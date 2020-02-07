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

package controlplane

import (
	"context"
	extensionswebhook "github.com/gardener/gardener-extensions/pkg/webhook"
	"github.com/gardener/gardener-extensions/pkg/webhook/controlplane"
	"github.com/gardener/gardener-extensions/pkg/webhook/controlplane/genericmutator"

	v1beta1constants "github.com/gardener/gardener/pkg/apis/core/v1beta1/constants"
	versionutils "github.com/gardener/gardener/pkg/utils/version"
	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// NewEnsurer creates a new controlplane ensurer.
func NewEnsurer(logger logr.Logger) genericmutator.Ensurer {
	return &ensurer{
		logger: logger.WithName("mock-controlplane-ensurer"),
	}
}

type ensurer struct {
	genericmutator.NoopEnsurer
	client client.Client
	logger logr.Logger
}

// InjectClient injects the given client into the ensurer.
func (e *ensurer) InjectClient(client client.Client) error {
	e.client = client
	return nil
}

// EnsureKubeAPIServerDeployment ensures that the kube-apiserver deployment conforms to the provider requirements.
func (e *ensurer) EnsureKubeAPIServerDeployment(ctx context.Context, ectx genericmutator.EnsurerContext, dep *appsv1.Deployment) error {
	template := &dep.Spec.Template
	ps := &template.Spec

	cluster, err := ectx.GetCluster(ctx)
	if err != nil {
		return err
	}

	if c := extensionswebhook.ContainerWithName(ps.Containers, "kube-apiserver"); c != nil {
		ensureVolumeMounts(c, cluster.Shoot.Spec.Kubernetes.Version)
	}
	ensureVolumes(ps, cluster.Shoot.Spec.Kubernetes.Version)
	return e.ensureChecksumAnnotations(ctx, &dep.Spec.Template, dep.Namespace)
}

// EnsureKubeControllerManagerDeployment ensures that the kube-controller-manager deployment conforms to the provider requirements.
func (e *ensurer) EnsureKubeControllerManagerDeployment(ctx context.Context, ectx genericmutator.EnsurerContext, dep *appsv1.Deployment) error {
	template := &dep.Spec.Template
	ps := &template.Spec

	cluster, err := ectx.GetCluster(ctx)
	if err != nil {
		return err
	}

	if c := extensionswebhook.ContainerWithName(ps.Containers, "kube-controller-manager"); c != nil {
		ensureVolumeMounts(c, cluster.Shoot.Spec.Kubernetes.Version)
	}
	ensureKubeControllerManagerAnnotations(template)
	ensureVolumes(ps, cluster.Shoot.Spec.Kubernetes.Version)
	return e.ensureChecksumAnnotations(ctx, &dep.Spec.Template, dep.Namespace)
}

func ensureKubeControllerManagerAnnotations(t *corev1.PodTemplateSpec) {
	t.Labels = extensionswebhook.EnsureAnnotationOrLabel(t.Labels, v1beta1constants.LabelNetworkPolicyToPublicNetworks, v1beta1constants.LabelNetworkPolicyAllowed)
	t.Labels = extensionswebhook.EnsureAnnotationOrLabel(t.Labels, v1beta1constants.LabelNetworkPolicyToPrivateNetworks, v1beta1constants.LabelNetworkPolicyAllowed)
	t.Labels = extensionswebhook.EnsureAnnotationOrLabel(t.Labels, v1beta1constants.LabelNetworkPolicyToBlockedCIDRs, v1beta1constants.LabelNetworkPolicyAllowed)
}

var (
	etcSSLName        = "etc-ssl"
	etcSSLVolumeMount = corev1.VolumeMount{
		Name:      etcSSLName,
		MountPath: "/etc/ssl",
		ReadOnly:  true,
	}
	etcSSLVolume = corev1.Volume{
		Name: etcSSLName,
		VolumeSource: corev1.VolumeSource{
			HostPath: &corev1.HostPathVolumeSource{
				Path: "/etc/ssl",
			},
		},
	}
)

func ensureVolumeMounts(c *corev1.Container, version string) {
	if mustMountEtcSSLFolder(version) {
		c.VolumeMounts = extensionswebhook.EnsureVolumeMountWithName(c.VolumeMounts, etcSSLVolumeMount)
	}
}

func ensureVolumes(ps *corev1.PodSpec, version string) {
	if mustMountEtcSSLFolder(version) {
		ps.Volumes = extensionswebhook.EnsureVolumeWithName(ps.Volumes, etcSSLVolume)
	}
}

// Beginning with 1.17 Gardener no longer uses the hyperkube image for the Kubernetes control plane components.
// The hyperkube image contained all the well-known root CAs, but the dedicated images don't. This is why we
// mount the /etc/ssl folder from the host here.
// TODO: This can be remove again once we have migrated to CSI.
func mustMountEtcSSLFolder(version string) bool {
	k8sVersionAtLeast117, err := versionutils.CompareVersions(version, ">=", "1.17")
	if err != nil {
		return false
	}
	return k8sVersionAtLeast117
}

func (e *ensurer) ensureChecksumAnnotations(ctx context.Context, template *corev1.PodTemplateSpec, namespace string) error {
	return controlplane.EnsureSecretChecksumAnnotation(ctx, template, e.client, namespace, v1beta1constants.SecretNameCloudProvider)
}
