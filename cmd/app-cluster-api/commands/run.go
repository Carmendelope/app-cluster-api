/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package commands

import (
    "github.com/nalej/app-cluster-api/internal/pkg/server"
    "github.com/rs/zerolog/log"
    "github.com/spf13/cobra"
)

var config = server.Config{}

var runCmd = &cobra.Command{
    Use:   "run",
    Short: "Run Cluster API",
    Long:  `Run Cluster API`,
    Run: func(cmd *cobra.Command, args []string) {
        SetupLogging()
        log.Info().Msg("Launching API!")
        server := server.NewService(config)
        server.Run()
    },
}

func init() {
    runCmd.Flags().IntVar(&config.Port, "port", 8281, "Port to launch the Public gRPC API")
    runCmd.Flags().StringVar(&config.DeploymentManagerAddress, "deploymentManagerAddress", "deployment-manager.nalej:5200", "deployment manager service address")
    runCmd.Flags().StringVar(&config.MusicianAddress, "musicianAddress", "musician.nalej:5100", "musician service address")

    rootCmd.AddCommand(runCmd)
}