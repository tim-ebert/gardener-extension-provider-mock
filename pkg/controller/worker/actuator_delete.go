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
	"github.com/gardener/gardener-extension-provider-mock/pkg/mock"
	"github.com/gardener/gardener-extensions/pkg/controller"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/apimachinery/pkg/util/wait"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

func (a *actuator) Delete(ctx context.Context, worker *extensionsv1alpha1.Worker, cluster *controller.Cluster) error {
	if err := extensionscontroller.DeleteManagedResource(ctx, a.Client(), worker.Namespace, mock.MockWorkerSecretName); err != nil {
		return fmt.Errorf("failed to delete managed resource for mock worker: %+v", err)
	}

	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	if err := extensionscontroller.WaitUntilManagedResourceDeleted(timeoutCtx, a.Client(), worker.Namespace, mock.MockWorkerSecretName); err != nil {
		return err
	}

	return waitUntilAllWorkersDeleted(ctx, a.Client(), 2*time.Second)
}

func waitUntilAllWorkersDeleted(ctx context.Context, c client.Client, interval time.Duration) error {
	workerSelector := labels.SelectorFromSet(map[string]string{
		"app":  "mock-worker",
		"role": "worker",
	})

	podList := &corev1.PodList{}
	return wait.PollUntil(interval, func() (done bool, err error) {
		if err := c.List(ctx, podList, client.MatchingLabelsSelector{Selector: workerSelector}); err != nil {
			return false, err
		}

		return len(podList.Items) == 0, nil
	}, ctx.Done())
}
