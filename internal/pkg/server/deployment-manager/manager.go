/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package deployment_manager

import (
    "github.com/nalej/grpc-deployment-manager-go"
    "context"
    "github.com/nalej/grpc-common-go"
)

// Manager to interact with the deployment manager
type Manager struct {
    DMClient grpc_deployment_manager_go.DeploymentManagerClient
}

// New manager
func NewManager(dmClient grpc_deployment_manager_go.DeploymentManagerClient) Manager {
    return Manager{DMClient: dmClient}
}

func (m *Manager) Execute(request *grpc_deployment_manager_go.DeploymentFragmentRequest) (*grpc_deployment_manager_go.DeploymentFragmentResponse, error) {
    return m.DMClient.Execute(context.Background(), request)
}

func (m *Manager) Undeploy(request *grpc_deployment_manager_go.UndeployRequest) (*grpc_common_go.Success, error) {
    return m.DMClient.Undeploy(context.Background(), request)
}