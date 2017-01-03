package main

import (
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var config = InitConfig()

	var serviceManager = InitServiceManager()
	var fileSystem = FileSystem{}
	var cmd = &CMD{}
	var projectSerialization = &ProjectSerialization{}
	var projectPaths = ProjectPaths{config}
	var projectLogic = &ProjectLogic{config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd}
	var daemonAction = DaemonAction{projectLogic}
	var projectActions = &ProjectActions{projectLogic}
	var cli = CLI{daemonAction, projectActions}

	cli.Init()
}
