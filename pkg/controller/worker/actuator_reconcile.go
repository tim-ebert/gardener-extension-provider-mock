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

package worker

import (
	"context"
	"fmt"
	mockv1alpha1 "github.com/gardener/gardener-extension-provider-mock/pkg/apis/mock/v1alpha1"
	"github.com/gardener/gardener-extension-provider-mock/pkg/imagevector"
	"github.com/gardener/gardener-extension-provider-mock/pkg/mock"
	"github.com/gardener/gardener-extensions/pkg/controller"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	"github.com/gardener/gardener-resource-manager/pkg/manager"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/gardener/gardener/pkg/operation/common"
	"github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/util/retry"
)

func (a *actuator) Reconcile(ctx context.Context, worker *extensionsv1alpha1.Worker, cluster *controller.Cluster) error {
	// Use v1beta1constants.SeedResourceManagerClass as mr class to deploy to seed

	var (
		workerConfig *mockv1alpha1.WorkerConfig
		err          error
	)

	poolSpec := worker.Spec.Pools[0]

	if poolSpec.ProviderConfig != nil {
		if _, _, err := a.Decoder().Decode(poolSpec.ProviderConfig.Raw, nil, workerConfig); err != nil {
			return fmt.Errorf("could not decode provider config: %+v", err)
		}
	}

	values, err := computeMockWorkerChartValues(worker, workerConfig, cluster)
	if err != nil {
		return err
	}
	release, err := a.ChartRenderer().Render(mock.MocknetChartPath, mock.MocknetReleaseName, metav1.NamespaceSystem, values)
	if err != nil {
		return err
	}

	if err := manager.NewSecret(a.Client()).
		WithKeyValues(map[string][]byte{mock.MocknetConfigKey: release.Manifest()}).
		WithNamespacedName(worker.Namespace, mock.MocknetConfigSecretName).
		Reconcile(ctx); err != nil {
		return err
	}

	if err := manager.
		NewManagedResource(a.Client()).
		WithNamespacedName(worker.Namespace, mock.MocknetConfigSecretName).
		WithSecretRefs([]corev1.LocalObjectReference{{Name: mock.MocknetConfigSecretName}}).
		WithInjectedLabels(map[string]string{common.ShootNoCleanup: "true"}).
		Reconcile(ctx)
		err != nil {
		return err
	}

	if err := a.updateWorkerStatus(ctx, worker); err != nil {
		return errors.Wrapf(err, "failed to update the machine deployments in worker status")
	}

	return nil
}

// computeMockWorkerChartValues computes the values for the mock-worker chart.
func computeMockWorkerChartValues(worker *extensionsv1alpha1.Worker, config *mockv1alpha1.WorkerConfig, cluster *extensionscontroller.Cluster) (map[string]interface{}, error) {
	values := map[string]interface{}{
		"images": map[string]interface{}{
			mock.HyperkubeImageName: imagevector.HyperkubeImage(cluster.Shoot.Spec.Kubernetes.Version),
			mock.DinDImageName:      imagevector.DinDImage(),
		},
		"nodeCount": 1, // only single node mock shoot supported for now
		"worker": map[string]interface{}{
			"pool": worker.Spec.Pools[0].Name,
		},
	}
	return values, nil
}

func (a *actuator) updateWorkerStatus(ctx context.Context, worker *extensionsv1alpha1.Worker) error {
	var statusMachineDeployments []extensionsv1alpha1.MachineDeployment

	return extensionscontroller.TryUpdateStatus(ctx, retry.DefaultBackoff, a.Client(), worker, func() error {
		worker.Status.MachineDeployments = statusMachineDeployments
		return nil
	})
}
