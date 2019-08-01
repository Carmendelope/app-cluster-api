/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package server

import (
    "github.com/nalej/derrors"
    "github.com/rs/zerolog/log"
    "github.com/nalej/app-cluster-api/version"
)

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
}


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

    return nil
}

func (conf *Config) Print() {
    log.Info().Str("app", version.AppVersion).Str("commit", version.Commit).Msg("Version")
    log.Info().Int("port", conf.Port).Msg("gRPC port")
    log.Info().Str("URL", conf.DeploymentManagerAddress).Msg("Deployment Manager Service")
    log.Info().Str("URL", conf.MusicianAddress).Msg("Musician Service")
    log.Info().Str("URL", conf.UnifiedLoggingAddress).Msg("Unified Logging Slave Service")
    log.Info().Str("URL", conf.MetricsCollectorAddress).Msg("Metrics Collector Service")
}

