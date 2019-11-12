/*
 * Copyright 2019 Nalej
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package cluster_watcher

import (
	"context"
	"github.com/nalej/grpc-cluster-watcher-go"
	"github.com/nalej/grpc-common-go"
)

// Manager to interact with the deployment manager
type Manager struct {
	CWClient grpc_cluster_watcher_go.ClusterWatcherSlaveClient
}

// New manager
func NewManager(cwClient grpc_cluster_watcher_go.ClusterWatcherSlaveClient) Manager {
	return Manager{CWClient: cwClient}
}

func (m *Manager) UpdateClusters(request *grpc_cluster_watcher_go.ListClustersWatchInfo) (*grpc_common_go.Success, error) {
	return m.CWClient.UpdateClusters(context.Background(), request)
}
