/*
 * Copyright (C) 2019 Nalej - All Rights Reserved
 */

package unified_logging

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-unified-logging-go"
	"github.com/rs/zerolog/log"
)

// Manager to interact with the unified logging slave service
type Manager struct {
	UnifiedLoggingClient grpc_unified_logging_go.SlaveClient
}

// New manager
func NewManager(unifiedLoggingClient grpc_unified_logging_go.SlaveClient) Manager {
	return Manager{UnifiedLoggingClient: unifiedLoggingClient}
}

func (m *Manager) Search(request *grpc_unified_logging_go.SearchRequest) (*grpc_unified_logging_go.LogResponse, error) {
	log.Debug().Interface("request", request).Msg("forward search request")
	return m.UnifiedLoggingClient.Search(context.Background(), request)
}

func (m *Manager) Expire(request *grpc_unified_logging_go.ExpirationRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("forward expiration request")
	return m.UnifiedLoggingClient.Expire(context.Background(), request)
}
