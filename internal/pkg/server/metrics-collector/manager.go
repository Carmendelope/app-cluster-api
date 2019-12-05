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
 */

package metrics_collector

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-monitoring-go"
	"github.com/rs/zerolog/log"
)

// Manager to interact with the unified logging slave service
type Manager struct {
	MetricsCollectorClient grpc_monitoring_go.MetricsCollectorClient
}

// New manager
func NewManager(mcClient grpc_monitoring_go.MetricsCollectorClient) Manager {
	return Manager{MetricsCollectorClient: mcClient}
}

// Retrieve a summary of high level cluster resource availability
func (m *Manager) GetClusterSummary(request *grpc_monitoring_go.ClusterSummaryRequest) (*grpc_monitoring_go.ClusterSummary, error) {
	log.Debug().Interface("request", request).Msg("forward GetClusterSummary request")
	return m.MetricsCollectorClient.GetClusterSummary(context.Background(), request)
}

// Retrieve statistics on cluster with respect to platform resources
func (m *Manager) GetClusterStats(request *grpc_monitoring_go.ClusterStatsRequest) (*grpc_monitoring_go.ClusterStats, error) {
	log.Debug().Interface("request", request).Msg("forward GetClusterStats request")
	return m.MetricsCollectorClient.GetClusterStats(context.Background(), request)
}

// Execute a query directly on the monitoring storage backend
func (m *Manager) Query(request *grpc_monitoring_go.QueryRequest) (*grpc_monitoring_go.QueryResponse, error) {
	log.Debug().Interface("request", request).Msg("forward Query request")
	return m.MetricsCollectorClient.Query(context.Background(), request)
}

// Execute a query directly on the monitoring storage backend
func (m *Manager) GetContainerStats(empty *grpc_common_go.Empty) (*grpc_monitoring_go.ContainerStatsResponse, error) {
	log.Debug().Msg("forward GetContainerStats request")
	return m.MetricsCollectorClient.GetContainerStats(context.Background(), empty)
}
