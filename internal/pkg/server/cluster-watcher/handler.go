/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package cluster_watcher

import (
	"context"
	"github.com/nalej/grpc-cluster-watcher-go"
	"github.com/nalej/grpc-common-go"
	"github.com/rs/zerolog/log"
)

// Handler structure for the conductor requests.
type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

// Request to score an application
func (h *Handler) UpdateClusters(context context.Context, request *grpc_cluster_watcher_go.ListClustersWatchInfo) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("update clusters")
	return h.Manager.UpdateClusters(request)
}
