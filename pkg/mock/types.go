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

package mock

import (
	"path/filepath"
)

const (
	// Name is the name of the Mock provider.
	Name = "provider-mock"

	// BackupSecretName is the name of the secret containing the credentials for storing the backups of Shoot clusters.
	BackupSecretName = "etcd-backup"

	// MocknetSecretName is the name of the secret used for the managed resource of networking mocknet
	MocknetSecretName = "extension-networking-mocknet"
	// MocknetReleaseName is the name of the Mocknet Release
	MocknetReleaseName = "mocknet"
	// MocknetConfigKey is the name of the key inside the mocknet config secret for storing the networking config.
	MocknetConfigKey = "mocknet.yaml"

	// MockWorkerSecretName is the name of the secret used for the managed resource for mock workers
	MockWorkerSecretName = "extension-worker-mock"
	// MockWorkerReleaseName is the name of the mock worker Release
	MockWorkerReleaseName = "mock-worker"
	// MockWorkerKey is the name of the key inside the mock worker secret for storing the rendered worker chart.
	MockWorkerKey = "mock-worker.yaml"
	// BootstrapTokenSecretName is the name of the secret used for the managed resource for bootstrap token for mock workers
	BootstrapTokenSecretName = "extension-worker-mock-bootstrap-token"

	// ImageNames
	CalicoCNIImageName                         = "calico-cni"
	CalicoNodeImageName                        = "calico-node"
	CalicoKubeControllersImageName             = "calico-kube-controllers"
	CalicoPodToDaemonFlexVolumeDriverImageName = "calico-podtodaemon-flex"
	HyperkubeImageName                         = "hyperkube"
	DinDImageName                              = "dind"
)

var (
	// ChartsPath is the path to the charts
	ChartsPath = "charts"
	// Interna1lChartsPath is the path to the internal charts
	InternalChartsPath = filepath.Join(ChartsPath, "internal")
	// MocknetChartPath is the path for internal Calico Chart
	MocknetChartPath = filepath.Join(InternalChartsPath, "mocknet")
	// MockWorkerChartPath is the path for internal Mock Worker Chart
	MockWorkerChartPath = filepath.Join(InternalChartsPath, "mock-worker")
)
