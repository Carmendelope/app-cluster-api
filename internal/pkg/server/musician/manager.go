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

package musician

import (
	"context"
	"github.com/nalej/grpc-conductor-go"
	"github.com/rs/zerolog/log"
)

// Manager to interact with the deployment manager
type Manager struct {
	MusicianClient grpc_conductor_go.MusicianClient
}

// New manager
func NewManager(musicianClient grpc_conductor_go.MusicianClient) Manager {
	return Manager{MusicianClient: musicianClient}
}

func (m *Manager) Score(request *grpc_conductor_go.ClusterScoreRequest) (*grpc_conductor_go.ClusterScoreResponse, error) {
	log.Debug().Interface("request", request).Msg("forward score request")
	return m.MusicianClient.Score(context.Background(), request)
}
