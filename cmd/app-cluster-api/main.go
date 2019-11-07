/*
 * Copyright (C) 2018 Nalej - All Rights Reserved
 */

package main

import (
	"github.com/nalej/app-cluster-api/cmd/app-cluster-api/commands"
	"github.com/nalej/app-cluster-api/version"
)

var MainVersion string

var MainCommit string

func main() {
	version.AppVersion = MainVersion
	version.Commit = MainCommit
	commands.Execute()
}
