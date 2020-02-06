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
	"github.com/gardener/gardener-resource-manager/pkg/manager"
	gardencorev1beta1 "github.com/gardener/gardener/pkg/apis/core/v1beta1"
	"github.com/gardener/gardener/pkg/operation/common"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/retry"

	mockv1alpha1 "github.com/gardener/gardener-extension-provider-mock/pkg/apis/mock/v1alpha1"
	"github.com/gardener/gardener-extension-provider-mock/pkg/imagevector"
	"github.com/gardener/gardener-extension-provider-mock/pkg/mock"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Reconcile implements Network.Actuator.
func (a *actuator) Reconcile(ctx context.Context, network *extensionsv1alpha1.Network, cluster *extensionscontroller.Cluster) error {
	var (
		networkConfig *mockv1alpha1.NetworkConfig
		err           error
	)

	if network.Spec.ProviderConfig != nil {
		if _, _, err := a.Decoder().Decode(network.Spec.ProviderConfig.Raw, nil, networkConfig); err != nil {
			return fmt.Errorf("could not decode provider config: %+v", err)
		}
	}

	values, err := computeMocknetChartValues(network, networkConfig)
	if err != nil {
		return err
	}
	release, err := a.ChartRenderer().Render(mock.MocknetChartPath, mock.MocknetReleaseName, metav1.NamespaceSystem, values)
	if err != nil {
		return err
	}

	if err := manager.NewSecret(a.Client()).
		WithKeyValues(map[string][]byte{mock.MocknetConfigKey: release.Manifest()}).
		WithNamespacedName(network.Namespace, mock.MocknetConfigSecretName).
		Reconcile(ctx); err != nil {
		return err
	}

	if err := manager.
		NewManagedResource(a.Client()).
		WithNamespacedName(network.Namespace, mock.MocknetConfigSecretName).
		WithSecretRefs([]corev1.LocalObjectReference{{Name: mock.MocknetConfigSecretName}}).
		WithInjectedLabels(map[string]string{common.ShootNoCleanup: "true"}).
		Reconcile(ctx)
		err != nil {
		return err
	}

	return a.updateProviderStatus(ctx, network, networkConfig)
}

// computeMocknetChartValues computes the values for the mocknet chart.
func computeMocknetChartValues(network *extensionsv1alpha1.Network, config *mockv1alpha1.NetworkConfig) (map[string]interface{}, error) {
	values := map[string]interface{}{
		"images": map[string]interface{}{
			mock.CalicoNodeImageName:                        imagevector.CalicoNodeImage(),
			mock.CalicoCNIImageName:                         imagevector.CalicoCNIImage(),
			mock.CalicoKubeControllersImageName:             imagevector.CalicoKubeControllersImage(),
			mock.CalicoPodToDaemonFlexVolumeDriverImageName: imagevector.CalicoFlexVolumeDriverImage(),
		},
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
