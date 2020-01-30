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

package infrastructure

import (
	"context"
	"fmt"
	mockapi "github.com/gardener/gardener-extension-provider-mock/pkg/apis/mock"
	mockv1alpha1 "github.com/gardener/gardener-extension-provider-mock/pkg/apis/mock/v1alpha1"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/util/retry"
)

func (a *actuator) reconcile(ctx context.Context, infrastructure *extensionsv1alpha1.Infrastructure, cluster *extensionscontroller.Cluster) error {
	infrastructureConfig := &mockapi.InfrastructureConfig{}
	if _, _, err := a.Decoder().Decode(infrastructure.Spec.ProviderConfig.Raw, nil, infrastructureConfig); err != nil {
		return fmt.Errorf("could not decode provider config: %+v", err)
	}

	return a.updateProviderStatus(ctx, infrastructure, infrastructureConfig)
}

func (a *actuator) updateProviderStatus(ctx context.Context, infrastructure *extensionsv1alpha1.Infrastructure, infrastructureConfig *mockapi.InfrastructureConfig) error {
	return extensionscontroller.TryUpdateStatus(ctx, retry.DefaultBackoff, a.Client(), infrastructure, func() error {
		infrastructure.Status.ProviderStatus = &runtime.RawExtension{
			Object: &mockv1alpha1.InfrastructureStatus{
				TypeMeta: metav1.TypeMeta{
					APIVersion: mockv1alpha1.SchemeGroupVersion.String(),
					Kind:       "InfrastructureStatus",
				},
			},
		}
		return nil
	})
}
