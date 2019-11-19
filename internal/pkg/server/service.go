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

package server

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"github.com/nalej/app-cluster-api/internal/pkg/server/cluster-watcher"
	"github.com/nalej/app-cluster-api/internal/pkg/server/deployment-manager"
	"github.com/nalej/app-cluster-api/internal/pkg/server/metrics-collector"
	"github.com/nalej/app-cluster-api/internal/pkg/server/musician"
	"github.com/nalej/app-cluster-api/internal/pkg/server/unified-logging"
	"github.com/nalej/derrors"
	"github.com/nalej/grpc-app-cluster-api-go"
	"github.com/nalej/grpc-cluster-watcher-go"
	"github.com/nalej/grpc-conductor-go"
	"github.com/nalej/grpc-deployment-manager-go"
	"github.com/nalej/grpc-monitoring-go"
	"github.com/nalej/grpc-unified-logging-go"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
	"io/ioutil"
	"net"
)

// Service structure with the configuration and the gRPC server.
type Service struct {
	Configuration Config
}

// Clients structure with the gRPC clients for remote services.
func NewService(conf Config) *Service {
	return &Service{
		conf,
	}
}

type Clients struct {
	DMClient               grpc_deployment_manager_go.DeploymentManagerClient
	DMNetworkClient        grpc_deployment_manager_go.DeploymentManagerNetworkClient
	MusicianClient         grpc_conductor_go.MusicianClient
	UnifiedLoggingClient   grpc_unified_logging_go.SlaveClient
	MetricsCollectorClient grpc_monitoring_go.MetricsCollectorClient
	ClusterWatcherClient   grpc_cluster_watcher_go.ClusterWatcherSlaveClient
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

	metricsCollectorConn, err := grpc.Dial(s.Configuration.MetricsCollectorAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with the metrics collector")
	}

	clusterWatcherConn, err := grpc.Dial(s.Configuration.ClusterWatcherAddress, grpc.WithInsecure())
	if err != nil {
		return nil, derrors.AsError(err, "cannot create connection with the cluster watcher")
	}

	dmClient := grpc_deployment_manager_go.NewDeploymentManagerClient(dmConn)
	dmNetworkClient := grpc_deployment_manager_go.NewDeploymentManagerNetworkClient(dmConn)
	musicianClient := grpc_conductor_go.NewMusicianClient(musicianConn)
	unifiedLoggingClient := grpc_unified_logging_go.NewSlaveClient(unifiedLoggingConn)
	metricsCollectorClient := grpc_monitoring_go.NewMetricsCollectorClient(metricsCollectorConn)
	clusterWatcherClient := grpc_cluster_watcher_go.NewClusterWatcherSlaveClient(clusterWatcherConn)

	return &Clients{
		DMClient:               dmClient,
		DMNetworkClient:        dmNetworkClient,
		MusicianClient:         musicianClient,
		UnifiedLoggingClient:   unifiedLoggingClient,
		MetricsCollectorClient: metricsCollectorClient,
		ClusterWatcherClient:   clusterWatcherClient,
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

	dmManager := deployment_manager.NewManager(clients.DMClient, clients.DMNetworkClient)
	dmHandler := deployment_manager.NewHandler(dmManager)

	musicianManager := musician.NewManager(clients.MusicianClient)
	musicianHandler := musician.NewHandler(musicianManager)

	ulManager := unified_logging.NewManager(clients.UnifiedLoggingClient)
	ulHandler := unified_logging.NewHandler(ulManager)

	mcManager := metrics_collector.NewManager(clients.MetricsCollectorClient)
	mcHandler := metrics_collector.NewHandler(mcManager)

	cwManager := cluster_watcher.NewManager(clients.ClusterWatcherClient)
	cwHandler := cluster_watcher.NewHandler(cwManager)
	// Create handlers
	creds, cErr := s.getSecureOptions()
	if cErr != nil{
		log.Fatal().Str("trace", cErr.DebugReport()).Msg("cannot build TLS options")
	}
	grpcServer := grpc.NewServer(grpc.Creds(creds), WithClientCertValidator(&s.Configuration))

	grpc_app_cluster_api_go.RegisterDeploymentManagerServer(grpcServer, dmHandler)
	grpc_app_cluster_api_go.RegisterMusicianServer(grpcServer, musicianHandler)
	grpc_app_cluster_api_go.RegisterUnifiedLoggingServer(grpcServer, ulHandler)
	grpc_app_cluster_api_go.RegisterMetricsCollectorServer(grpcServer, mcHandler)
	grpc_app_cluster_api_go.RegisterClusterWatcherSlaveServer(grpcServer, cwHandler)

	reflection.Register(grpcServer)
	log.Info().Int("port", s.Configuration.Port).Msg("Launching gRPC server")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal().Errs("failed to serve: %v", []error{err})
	}
	return nil
}

func (s * Service) getSecureOptions() (credentials.TransportCredentials, derrors.Error) {
	rootCAs := x509.NewCertPool()
	tlsConfig := &tls.Config{}

	// Load CA to validate client certificate
	log.Debug().Str("caCertPath", s.Configuration.CACertPath).Msg("loading server certificate")
	serverCert, err := ioutil.ReadFile(s.Configuration.CACertPath)
	if err != nil {
		return nil, derrors.NewInternalError("Error loading server certificate")
	}
	added := rootCAs.AppendCertsFromPEM(serverCert)
	if !added {
		return nil, derrors.NewInternalError("cannot add server certificate to the pool")
	}
	tlsConfig.RootCAs = rootCAs
	// Client certificate for app-cluster-api
	log.Debug().Str("clientCertPath", s.Configuration.ClientCertPath).Msg("loading client certificate")
	clientCert, err := tls.LoadX509KeyPair(fmt.Sprintf("%s/tls.crt", s.Configuration.ClientCertPath), fmt.Sprintf("%s/tls.key", s.Configuration.ClientCertPath))
	if err != nil {
		log.Error().Str("error", err.Error()).Msg("Error loading client certificate")
		return nil, derrors.NewInternalError("Error loading client certificate")
	}

	tlsConfig.Certificates = []tls.Certificate{clientCert}
	tlsConfig.BuildNameToCertificate()
	tlsConfig.VerifyPeerCertificate = s.verifyManagementCluster
	creds := credentials.NewTLS(tlsConfig)
	log.Debug().Interface("creds", creds.Info()).Msg("Secure credentials")
	return creds, nil
}

func (s * Service) verifyManagementCluster(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	log.Info().Msg("verifyManagementCluster")
	log.Info().Int("numRawCerts", len(rawCerts)).Int("numVerifiedChains", len(verifiedChains)).Msg("parameters")
	for _, vc := range verifiedChains{
		log.Info().Interface("vc", vc).Msg("chain element")
	}
	return nil
}
