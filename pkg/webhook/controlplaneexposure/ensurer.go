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

package controlplaneexposure

import (
	"context"
	"github.com/gardener/gardener-extension-provider-mock/pkg/apis/config"
	extensionswebhook "github.com/gardener/gardener-extensions/pkg/webhook"
	"github.com/gardener/gardener-extensions/pkg/webhook/controlplane"
	"github.com/gardener/gardener-extensions/pkg/webhook/controlplane/genericmutator"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// NewEnsurer creates a new controlplaneexposure ensurer.
func NewEnsurer(etcdStorage *config.ETCDStorage, logger logr.Logger) genericmutator.Ensurer {
	return &ensurer{
		etcdStorage: etcdStorage,
		logger:      logger.WithName("mock-controlplaneexposure-ensurer"),
	}
}

type ensurer struct {
	genericmutator.NoopEnsurer
	etcdStorage *config.ETCDStorage
	logger      logr.Logger
}

// EnsureKubeAPIServerService ensures that the kube-apiserver service conforms to the provider requirements.
func (e *ensurer) EnsureKubeAPIServerService(ctx context.Context, ectx genericmutator.EnsurerContext, svc *corev1.Service) error {
	return nil
}

// EnsureKubeAPIServerDeployment ensures that the kube-apiserver deployment conforms to the provider requirements.
func (e *ensurer) EnsureKubeAPIServerDeployment(ctx context.Context, ectx genericmutator.EnsurerContext, dep *appsv1.Deployment) error {
	return nil
}

// EnsureETCDStatefulSet ensures that the etcd stateful sets conform to the provider requirements.
func (e *ensurer) EnsureETCDStatefulSet(ctx context.Context, ectx genericmutator.EnsurerContext, ss *appsv1.StatefulSet) error {
	e.ensureVolumeClaimTemplates(&ss.Spec, ss.Name)
	return nil
}

func (e *ensurer) ensureVolumeClaimTemplates(spec *appsv1.StatefulSetSpec, name string) {
	t := e.getVolumeClaimTemplate(name)
	spec.VolumeClaimTemplates = extensionswebhook.EnsurePVCWithName(spec.VolumeClaimTemplates, *t)
}

func (e *ensurer) getVolumeClaimTemplate(name string) *corev1.PersistentVolumeClaim {
	var (
		etcdStorage             config.ETCDStorage
		volumeClaimTemplateName = name
		defaultClass            = "default"
	)

	if name == v1beta1constants.ETCDMain {
		// TODO storage class for mock provider
		etcdStorage = *e.etcdStorage
		etcdStorage.ClassName = &defaultClass
		volumeClaimTemplateName = controlplane.EtcdMainVolumeClaimTemplateName
	}

	return controlplane.GetETCDVolumeClaimTemplate(volumeClaimTemplateName, etcdStorage.ClassName, etcdStorage.Capacity)
}
