/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package infrastructure_monitor

import (
    "context"
    "github.com/rs/zerolog/log"
    grpc "github.com/nalej/grpc-infrastructure-monitor-go"
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
