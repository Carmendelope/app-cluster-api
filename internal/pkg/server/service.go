/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package server

import (
    "github.com/nalej/grpc-utils/pkg/tools"
    "github.com/nalej/grpc-conductor-go"
    "github.com/nalej/grpc-deployment-manager-go"
    "github.com/nalej/grpc-infrastructure-monitor-go"
    "github.com/nalej/grpc-unified-logging-go"
    "github.com/nalej/derrors"
    "google.golang.org/grpc"
    "net"
    "fmt"
    "github.com/nalej/app-cluster-api/internal/pkg/server/deployment-manager"
    "github.com/nalej/app-cluster-api/internal/pkg/server/musician"
    "github.com/nalej/app-cluster-api/internal/pkg/server/unified-logging"
    "github.com/nalej/app-cluster-api/internal/pkg/server/infrastructure-monitor"
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
    UnifiedLoggingClient grpc_unified_logging_go.SlaveClient
    InfrastructureMonitorClient grpc_infrastructure_monitor_go.SlaveClient
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

    unifiedLoggingConn, err := grpc.Dial(s.Configuration.UnifiedLoggingAddress, grpc.WithInsecure())
    if err != nil {
        return nil, derrors.AsError(err, "cannot create connection with the unified logging slave")
    }

    infrastructureMonitorConn, err := grpc.Dial(s.Configuration.InfrastructureMonitorAddress, grpc.WithInsecure())
    if err != nil {
        return nil, derrors.AsError(err, "cannot create connection with the infrastructure monitor slave")
    }

    dmClient := grpc_deployment_manager_go.NewDeploymentManagerClient(dmConn)
    musicianClient := grpc_conductor_go.NewMusicianClient(musicianConn)
    unifiedLoggingClient := grpc_unified_logging_go.NewSlaveClient(unifiedLoggingConn)
    infrastructureMonitorClient := grpc_infrastructure_monitor_go.NewSlaveClient(infrastructureMonitorConn)

    return &Clients{
        DMClient: dmClient,
        MusicianClient: musicianClient,
        UnifiedLoggingClient: unifiedLoggingClient,
        InfrastructureMonitorClient: infrastructureMonitorClient,
    }, nil
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

    ulManager := unified_logging.NewManager(clients.UnifiedLoggingClient)
    ulHandler := unified_logging.NewHandler(ulManager)

    imManager := infrastructure_monitor.NewManager(clients.InfrastructureMonitorClient)
    imHandler := infrastructure_monitor.NewHandler(imManager)
    // Create handlers
    grpcServer := grpc.NewServer()

    grpc_app_cluster_api_go.RegisterDeploymentManagerServer(grpcServer, dmHandler)
    grpc_app_cluster_api_go.RegisterMusicianServer(grpcServer, musicianHandler)
    grpc_app_cluster_api_go.RegisterUnifiedLoggingServer(grpcServer, ulHandler)
    grpc_app_cluster_api_go.RegisterInfrastructureMonitorServer(grpcServer, imHandler)

    reflection.Register(grpcServer)
    log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatal().Errs("failed to serve: %v", []error{err})
    }
    return nil
}
