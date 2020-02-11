// Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package network

import (
	"context"
	"fmt"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/gardener/gardener/pkg/utils/chart"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/retry"

	mockv1alpha1 "github.com/gardener/gardener-extension-provider-mock/pkg/apis/mock/v1alpha1"
	"github.com/gardener/gardener-extension-provider-mock/pkg/imagevector"
	"github.com/gardener/gardener-extension-provider-mock/pkg/mock"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var mocknetChart = &chart.Chart{
	Name: "mocknet",
	Path: mock.MocknetChartPath,
	Images: []string{
		mock.CalicoCNIImageName,
		mock.CalicoNodeImageName,
		mock.CalicoKubeControllersImageName,
		mock.CalicoPodToDaemonFlexVolumeDriverImageName,
	},
}

// Reconcile implements Network.Actuator.
func (a *actuator) Reconcile(ctx context.Context, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster) error {
	networkConfig := &mockv1alpha1.NetworkConfig{}
	if network.Spec.ProviderConfig != nil {
		if _, _, err := a.Decoder().Decode(network.Spec.ProviderConfig.Raw, nil, networkConfig); err != nil {
			return fmt.Errorf("could not decode provider config: %+v", err)
		}
	}

	values, err := computeMocknetChartValues(network, networkConfig)
	if err != nil {
		return err
	}

	if err := extensionscontroller.RenderChartAndCreateManagedResource(ctx, network.Namespace, mock.MocknetSecretName,
		a.Client(), a.ChartRenderer(), mocknetChart, values, imagevector.ImageVector(), metav1.NamespaceSystem,
		cluster.Shoot.Spec.Kubernetes.Version, true, false); err != nil {
		return fmt.Errorf("failed to deploy mocknet: %+v", err)
	}

	return a.updateProviderStatus(ctx, network, networkConfig)
}

// computeMocknetChartValues computes the values for the mocknet chart.
func computeMocknetChartValues(network *extensionsv1alpha1.Network, config *mockv1alpha1.NetworkConfig) (map[string]interface{}, error) {
	values := map[string]interface{}{
		"global": map[string]string{
			"podCIDR": network.Spec.PodCIDR,
		},
		"config": map[string]interface{}{
			"veth_mtu": 1440,
			"backend":  "bird",
			"ipam": map[string]interface{}{
				"type":   "host-local",
				"subnet": "usePodCidr",
			},
			"kubeControllers": map[string]interface{}{
				"enabled": true,
			},
			"ipv4": map[string]interface{}{
				"pool": "ipip",
				"mode": "Always",
			},
			"felix": map[string]interface{}{
				"ipinip": map[string]interface{}{
					"enabled": "true",
				},
			},
		},
	}
	return values, nil
}

func (a *actuator) updateProviderStatus(ctx context.Context, network *extensionsv1alpha1.Network, config *mockv1alpha1.NetworkConfig) error {
	status, err := a.ComputeNetworkStatus(config)
	if err != nil {
		return err
	}

	return extensionscontroller.TryUpdateStatus(ctx, retry.DefaultBackoff, a.Client(), network, func() error {
		network.Status.ProviderStatus = &runtime.RawExtension{Object: status}
		network.Status.LastOperation = extensionscontroller.LastOperation(gardencorev1beta1.LastOperationTypeReconcile,
			gardencorev1beta1.LastOperationStateSucceeded,
			100,
			"Mocknet was configured successfully")
		return nil
	})
}

func (a *actuator) ComputeNetworkStatus(networkConfig *mockv1alpha1.NetworkConfig) (*mockv1alpha1.NetworkStatus, error) {
	var (
		status = &mockv1alpha1.NetworkStatus{
			TypeMeta: StatusTypeMeta,
		}
	)

	return status, nil
}
