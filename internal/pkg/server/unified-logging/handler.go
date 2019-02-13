/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package unified_logging

import (
    "context"
    "github.com/rs/zerolog/log"
    "github.com/nalej/grpc-unified-logging-go"
    "github.com/nalej/grpc-common-go"
)

type Handler struct {
    Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
    return &Handler{manager}
}

func (h *Handler) Search(ctx context.Context, request *grpc_unified_logging_go.SearchRequest) (*grpc_unified_logging_go.LogResponse, error) {
    log.Debug().Interface("request", request).Msg("search")
    return h.Manager.Search(request)
}

func (h *Handler) Expire(ctx context.Context, request *grpc_unified_logging_go.ExpirationRequest) (*grpc_common_go.Success, error) {
    log.Debug().Interface("request", request).Msg("expire")
    return h.Manager.Expire(request)
}

