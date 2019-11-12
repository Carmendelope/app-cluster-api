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

package deployment_manager

import (
	"context"
	"github.com/nalej/grpc-common-go"
	"github.com/nalej/grpc-deployment-manager-go"
	"github.com/rs/zerolog/log"
)

// Manager to interact with the deployment manager
type Manager struct {
	DMClient        grpc_deployment_manager_go.DeploymentManagerClient
	DMNetworkClient grpc_deployment_manager_go.DeploymentManagerNetworkClient
}

// New manager
func NewManager(dmClient grpc_deployment_manager_go.DeploymentManagerClient, dmNetworkClient grpc_deployment_manager_go.DeploymentManagerNetworkClient) Manager {
	return Manager{DMClient: dmClient, DMNetworkClient: dmNetworkClient}
}

func (m *Manager) Execute(request *grpc_deployment_manager_go.DeploymentFragmentRequest) (*grpc_deployment_manager_go.DeploymentFragmentResponse, error) {
	log.Debug().Interface("request", request).Msg("forward execute request")
	return m.DMClient.Execute(context.Background(), request)
}

func (m *Manager) Undeploy(request *grpc_deployment_manager_go.UndeployRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("forward undeploy request")
	return m.DMClient.Undeploy(context.Background(), request)
}

func (m *Manager) UndeployFragment(request *grpc_deployment_manager_go.UndeployFragmentRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("forward undeploy fragment request")
	return m.DMClient.UndeployFragment(context.Background(), request)
}

func (m *Manager) SetServiceRoute(request *grpc_deployment_manager_go.ServiceRoute) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("forward set service route")
	return m.DMNetworkClient.SetServiceRoute(context.Background(), request)

}

// JoinZTNetwork
func (m *Manager) JoinZTNetwork(request *grpc_deployment_manager_go.JoinZTNetworkRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("forward join ZT-Network")
	return m.DMNetworkClient.JoinZTNetwork(context.Background(), request)
}

func (m *Manager) LeaveZTNetwork(request *grpc_deployment_manager_go.LeaveZTNetworkRequest) (*grpc_common_go.Success, error) {
	log.Debug().Interface("request", request).Msg("forward leave ZT-Network")
	return m.DMNetworkClient.LeaveZTNetwork(context.Background(), request)
}
