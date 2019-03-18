/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package infrastructure_monitor

import (
    "context"
    grpc "github.com/nalej/grpc-infrastructure-monitor-go"
    "github.com/rs/zerolog/log"
)

// Manager to interact with the unified logging slave service
type Manager struct {
    InfrastructureMonitorClient grpc.SlaveClient
}

// New manager
func NewManager(imClient grpc.SlaveClient) Manager {
    return Manager{InfrastructureMonitorClient: imClient}
}

// Retrieve a summary of high level cluster resource availability
func (m *Manager) GetClusterSummary(request *grpc.ClusterSummaryRequest) (*grpc.ClusterSummary, error) {
    log.Debug().Interface("request", request).Msg("forward GetClusterSummary request")
    return m.InfrastructureMonitorClient.GetClusterSummary(context.Background(), request)
}

// Retrieve statistics on cluster with respect to platform resources
func (m *Manager) GetClusterStats(request *grpc.ClusterStatsRequest) (*grpc.ClusterStats, error) {
    log.Debug().Interface("request", request).Msg("forward GetClusterStats request")
    return m.InfrastructureMonitorClient.GetClusterStats(context.Background(), request)
}

// Execute a query directly on the monitoring storage backend
func (m *Manager) Query(request *grpc.QueryRequest) (*grpc.QueryResponse, error) {
    log.Debug().Interface("request", request).Msg("forward Query request")
    return m.InfrastructureMonitorClient.Query(context.Background(), request)
}
