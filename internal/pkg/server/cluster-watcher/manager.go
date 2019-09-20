/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
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
