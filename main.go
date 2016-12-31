package main

import (
	"math/rand"
	"runtime"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var config = Config{
		OS:        runtime.GOOS,
		PathSep:   "/",
		Workspace: "/srv/molly",
	}

	var projectPaths = ProjectPaths{config}
	var projectLogic = ProjectLogic{config, projectPaths}
	var daemonAction = DaemonAction{projectLogic}
	var projectActions = ProjectActions{projectLogic}
	var cli = CLI{daemonAction, projectActions}

	cli.Init()
}
