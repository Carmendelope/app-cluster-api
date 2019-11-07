/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package musician

import (
	"context"
	"github.com/nalej/grpc-conductor-go"
	"github.com/rs/zerolog/log"
)

type Handler struct {
	Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
	return &Handler{manager}
}

// Request to score an application
func (h *Handler) Score(ctx context.Context, request *grpc_conductor_go.ClusterScoreRequest) (*grpc_conductor_go.ClusterScoreResponse, error) {
	log.Debug().Interface("request", request).Msg("score")
	return h.Manager.Score(request)
}
