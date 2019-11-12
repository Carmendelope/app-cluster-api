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
