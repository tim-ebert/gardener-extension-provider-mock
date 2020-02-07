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

package controlplanebackup

import (
	"context"
	"fmt"
	"github.com/gardener/gardener-extension-provider-mock/pkg/apis/config"
	"github.com/gardener/gardener-extension-provider-mock/pkg/mock"
	extensionscontroller "github.com/gardener/gardener-extensions/pkg/controller"
	extensionswebhook "github.com/gardener/gardener-extensions/pkg/webhook"
	"github.com/gardener/gardener-extensions/pkg/webhook/controlplane"
	"github.com/gardener/gardener-extensions/pkg/webhook/controlplane/genericmutator"
	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	"github.com/gardener/gardener/pkg/utils/imagevector"
	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewEnsurer creates a new controlplaneexposure ensurer.
func NewEnsurer(etcdBackup *config.ETCDBackup, imageVector imagevector.ImageVector, logger logr.Logger) genericmutator.Ensurer {
	return &ensurer{
		etcdBackup:  etcdBackup,
		imageVector: imageVector,
		logger:      logger.WithName("mock-controlplanebackup-ensurer"),
	}
}

type ensurer struct {
	genericmutator.NoopEnsurer
	etcdBackup  *config.ETCDBackup
	imageVector imagevector.ImageVector
	client      client.Client
	logger      logr.Logger
}

// InjectClient injects the given client into the ensurer.
func (e *ensurer) InjectClient(client client.Client) error {
	e.client = client
	return nil
}

// EnsureETCDStatefulSet ensures that the etcd stateful sets conform to the provider requirements.
func (e *ensurer) EnsureETCDStatefulSet(ctx context.Context, ectx genericmutator.EnsurerContext, ss *appsv1.StatefulSet) error {
	cluster, err := ectx.GetCluster(ctx)
	if err != nil {
		return err
	}
	if err := e.ensureContainers(&ss.Spec.Template.Spec, ss.Name, cluster); err != nil {
		return err
	}
	return e.ensureChecksumAnnotations(ctx, &ss.Spec.Template, ss.Namespace, ss.Name, cluster.Seed.Spec.Backup != nil)
}

func (e *ensurer) ensureContainers(ps *corev1.PodSpec, name string, cluster *extensionscontroller.Cluster) error {
	etcdContainer := extensionswebhook.ContainerWithName(ps.Containers, "etcd")
	c, err := e.ensureETCDContainer(etcdContainer, name)
	if err != nil {
		return err
	}
	ps.Containers = extensionswebhook.EnsureContainerWithName(ps.Containers, *c)
	return nil
}

func (e *ensurer) ensureChecksumAnnotations(ctx context.Context, template *corev1.PodTemplateSpec, namespace, name string, backupConfigured bool) error {
	if name == v1beta1constants.ETCDMain && backupConfigured {
		return controlplane.EnsureSecretChecksumAnnotation(ctx, template, e.client, namespace, mock.BackupSecretName)
	}
	return nil
}

func (e *ensurer) ensureETCDContainer(c *corev1.Container, name string) (*corev1.Container, error) {
	if c == nil {
		return nil, fmt.Errorf("could not find etcd container in pod spec")
	}

	c.Command = []string{
		"etcd", "--config-file", "/bootstrap/etcd.conf.yml",
	}

	c.ReadinessProbe.HTTPGet = nil
	c.ReadinessProbe.Exec = &corev1.ExecAction{Command: []string{
		"/bin/sh",
		"-ec",
		"ETCDCTL_API=3",
		"etcdctl",
		"--cert=/var/etcd/ssl/client/tls.crt",
		"--key=/var/etcd/ssl/client/tls.key",
		"--cacert=/var/etcd/ssl/ca/ca.crt",
		fmt.Sprintf("--endpoints=https://%s-0:%d", name, 2379),
		"endpoint",
		"health",
	}}

	return c, nil
}
