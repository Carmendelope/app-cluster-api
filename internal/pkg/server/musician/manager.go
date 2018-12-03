/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package musician

import (
    "github.com/nalej/grpc-conductor-go"
    "context"
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
    return m.MusicianClient.Score(context.Background(), request)
}