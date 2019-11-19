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
	"github.com/nalej/app-cluster-api/version"
	"github.com/nalej/derrors"
	"github.com/rs/zerolog/log"
)

// Config structure with the configuration parameters of the application.
type Config struct {
	// Port where the gRPC API service will listen  to incoming requests
	Port int
	// DeploymentManager address
	DeploymentManagerAddress string
	// Musician address
	MusicianAddress string
	// Unified Logging Slave address
	UnifiedLoggingAddress string
	// Metrics Collector address
	MetricsCollectorAddress string
	// Cluster watcher address
	ClusterWatcherAddress string
	// Path for the certificate of the CA
	CACertPath string
	// client certificate path to use for validation
	ClientCertPath string
	// ManagementPublicHost with the public host name of the management cluster.
	ManagementPublicHost string
}

// Validate checks the configuration options returning an error if any mandatory value is missing.
func (conf *Config) Validate() derrors.Error {
	if conf.Port <= 0 {
		return derrors.NewInvalidArgumentError("ports must be valid")
	}
	if conf.MusicianAddress == "" {
		return derrors.NewInvalidArgumentError("musicianAddress invalid")
	}
	if conf.DeploymentManagerAddress == "" {
		return derrors.NewInvalidArgumentError("deploymentManagerAddress invalid")
	}
	if conf.UnifiedLoggingAddress == "" {
		return derrors.NewInvalidArgumentError("unifiedLoggingAddress invalid")
	}
	if conf.MetricsCollectorAddress == "" {
		return derrors.NewInvalidArgumentError("metricsCollectorAddress invalid")
	}
	if conf.ClusterWatcherAddress == "" {
		return derrors.NewInvalidArgumentError("clusterWatcherAddress invalid")
	}
	if conf.CACertPath == "" {
		return derrors.NewInvalidArgumentError("caCertPath must be set")
	}
	if conf.ClientCertPath == "" {
		return derrors.NewInvalidArgumentError("clientCertPath must be set")
	}
	if conf.ManagementPublicHost == "" {
		return derrors.NewInvalidArgumentError("managementPublicHost must be set")
	}
	return nil
}

// Print the configuration values to the log.
func (conf *Config) Print() {
	log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
	log.Info().Int("port", conf.Port).Str("clientCertPath", conf.ClientCertPath).Str("caCertPath", conf.CACertPath).Msg("gRPC port")
	log.Info().Str("URL", conf.DeploymentManagerAddress).Msg("Deployment Manager Service")
	log.Info().Str("URL", conf.MusicianAddress).Msg("Musician Service")
	log.Info().Str("URL", conf.UnifiedLoggingAddress).Msg("Unified Logging Slave Service")
	log.Info().Str("URL", conf.MetricsCollectorAddress).Msg("Metrics Collector Service")
	log.Info().Str("URL", conf.ClusterWatcherAddress).Msg("Cluster watcher service")
	log.Info().Str("URL", conf.ManagementPublicHost).Msg("Management cluster")
}
