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
	"github.com/gardener/gardener-extensions/pkg/controller"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	"github.com/pkg/errors"
	"k8s.io/client-go/util/retry"
)

func (a *actuator) Reconcile(ctx context.Context, worker *extensionsv1alpha1.Worker, cluster *controller.Cluster) error {
	if err := a.updateWorkerStatus(ctx, worker); err != nil {
		return errors.Wrapf(err, "failed to update the machine deployments in worker status")
	}

	return nil
}

func (a *actuator) updateWorkerStatus(ctx context.Context, worker *extensionsv1alpha1.Worker) error {
	var statusMachineDeployments []extensionsv1alpha1.MachineDeployment

	return extensionscontroller.TryUpdateStatus(ctx, retry.DefaultBackoff, a.Client(), worker, func() error {
		worker.Status.MachineDeployments = statusMachineDeployments
		return nil
	})
}
