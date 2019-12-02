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

package metrics_collector

import (
	"context"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-common-go"
	grpc "github.com/nalej/grpc-monitoring-go"
	"github.com/nalej/grpc-utils/pkg/conversions"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

// Retrieve a summary of high level cluster resource availability
func (h *Handler) GetClusterSummary(ctx context.Context, request *grpc.ClusterSummaryRequest) (*grpc.ClusterSummary, error) {
	log.Debug().Interface("request", request).Msg("GetClusterSummary")
	return h.Manager.GetClusterSummary(request)
}

// Retrieve statistics on cluster with respect to platform resources
func (h *Handler) GetClusterStats(ctx context.Context, request *grpc.ClusterStatsRequest) (*grpc.ClusterStats, error) {
	log.Debug().Interface("request", request).Msg("GetClusterStats")
	return h.Manager.GetClusterStats(request)
}

// Execute a query directly on the monitoring storage backend
func (h *Handler) Query(ctx context.Context, request *grpc.QueryRequest) (*grpc.QueryResponse, error) {
	log.Debug().Interface("request", request).Msg("Query")
	return h.Manager.Query(request)
}

func (h *Handler) GetContainerStats(ctx context.Context, in *grpc_common_go.Empty) (*grpc.ContainerStatsResponse, error) {
	return nil, conversions.ToGRPCError(derrors.NewUnimplementedError("not implemented yet"))
}
