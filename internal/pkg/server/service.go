/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package server

import (
    "github.com/nalej/grpc-utils/pkg/tools"
    "github.com/nalej/grpc-deployment-manager-go"
    "github.com/nalej/grpc-conductor-go"
    "github.com/nalej/derrors"
    "google.golang.org/grpc"
    "net"
    "fmt"
    "github.com/nalej/app-cluster-api/internal/pkg/server/deployment-manager"
    "github.com/nalej/app-cluster-api/internal/pkg/server/musician"
    "github.com/rs/zerolog/log"
    "github.com/nalej/grpc-app-cluster-api-go"
    "google.golang.org/grpc/reflection"
)

// Service structure with the configuration and the gRPC server.
type Service struct {
    Configuration Config
    Server        *tools.GenericGRPCServer
}

// Clients structure with the gRPC clients for remote services.
func NewService(conf Config) *Service {
    return &Service {
        conf,
        tools.NewGenericGRPCServer(uint32(conf.Port)),
    }
}

type Clients struct {
    DMClient grpc_deployment_manager_go.DeploymentManagerClient
    MusicianClient grpc_conductor_go.MusicianClient
}

func (s *Service) GetClients() (*Clients, derrors.Error) {
    dmConn, err := grpc.Dial(s.Configuration.DeploymentManagerAddress, grpc.WithInsecure())
    if err != nil {
        return nil, derrors.AsError(err, "cannot create connection with the deployment manager")
    }

    musicianConn, err := grpc.Dial(s.Configuration.MusicianAddress, grpc.WithInsecure())
    if err != nil {
        return nil, derrors.AsError(err, "cannot create connection with the musician")
    }

    dmClient := grpc_deployment_manager_go.NewDeploymentManagerClient(dmConn)
    musicianClient := grpc_conductor_go.NewMusicianClient(musicianConn)

    return &Clients{DMClient: dmClient, MusicianClient: musicianClient}, nil
}

// Run the service, launch the service handler
func (s *Service) Run() error {
    cErr := s.Configuration.Validate()
    if cErr != nil {
        log.Fatal().Str("err", cErr.DebugReport()).Msg("invalid configuration")
    }
    s.Configuration.Print()

    clients, cErr := s.GetClients()
    if cErr != nil {
        log.Fatal().Str("err", cErr.DebugReport()).Msg("Cannot create clients")
    }

    lis, err := net.Listen("tcp", fmt.Sprintf(":%d", s.Configuration.Port))
    if err != nil {
        log.Fatal().Errs("failed to listen: %v", []error{err})
    }

    dmManager := deployment_manager.NewManager(clients.DMClient)
    dmHandler := deployment_manager.NewHandler(dmManager)

    musicianManager := musician.NewManager(clients.MusicianClient)
    musicianHandler := musician.NewHandler(musicianManager)

    // Create handlers
    grpcServer := grpc.NewServer()

    grpc_app_cluster_api_go.RegisterDeploymentManagerServer(grpcServer, dmHandler)
    grpc_app_cluster_api_go.RegisterMusicianServer(grpcServer, musicianHandler)

    reflection.Register(grpcServer)
    log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatal().Errs("failed to serve: %v", []error{err})
    }
    return nil
}
