/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package deployment_manager

import (
    "context"
    "github.com/rs/zerolog/log"
    "github.com/nalej/grpc-deployment-manager-go"
    "github.com/nalej/grpc-common-go"
)

// Handler structure for the conductor requests.
type Handler struct {
    Manager Manager
}

// NewHandler creates a new Handler with a linked manager.
func NewHandler(manager Manager) *Handler {
    return &Handler{manager}
}

func (h *Handler) Execute(context context.Context, request *grpc_deployment_manager_go.DeploymentFragmentRequest) (*grpc_deployment_manager_go.DeploymentFragmentResponse, error) {
    log.Debug().Interface("request", request).Msg("execute deployment fragment")
    return h.Manager.Execute(request)
}

func (h* Handler) Undeploy (context context.Context, request *grpc_deployment_manager_go.UndeployRequest) (*grpc_common_go.Success, error) {
    log.Debug().Interface("request", request).Msg("undeploy")
    return h.Manager.Undeploy(request)
}

func (h* Handler) UndeployFragment (context context.Context, request *grpc_deployment_manager_go.UndeployFragmentRequest) (*grpc_common_go.Success, error) {
    log.Debug().Interface("request", request).Msg("undeploy fragment")
    return h.Manager.UndeployFragment(request)
}


func (h* Handler) SetServiceRoute(context context.Context, request *grpc_deployment_manager_go.ServiceRoute) (*grpc_common_go.Success, error) {
    log.Debug().Interface("request", request).Msg("set service route")
    return h.Manager.SetServiceRoute(request)
}

// JoinZTNetwork message to Request a zt-agent to join into a new Network
func (h* Handler) JoinZTNetwork(_ context.Context, request *grpc_deployment_manager_go.JoinZTNetworkRequest) (*grpc_common_go.Success, error){
    log.Debug().Interface("request", request).Msg("join ZT-Network")
    return h.Manager.JoinZTNetwork(request)
}

func (h* Handler) LeaveZTNetwork(_ context.Context, request *grpc_deployment_manager_go.LeaveZTNetworkRequest) (*grpc_common_go.Success, error){
    log.Debug().Interface("request", request).Msg("leave ZT-Network")
    return h.Manager.LeaveZTNetwork(request)
}