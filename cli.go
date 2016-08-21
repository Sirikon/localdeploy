package main

import (
	"github.com/urfave/cli"
	"os"
)

func InitCLI() {
	app := cli.NewApp()

	app.Name = "molly"
	app.HelpName = "molly"
	app.Usage = "Minimalistic local deployment tool"
	app.Version = "0.0.1"
	app.Commands = []cli.Command{
		{
			Name: "daemon",
			Usage: "Starts the application as a daemon",
			Action: ActionDaemon,
		},
		{
			Name: "project",
			Usage: "Project management",
			Subcommands: []cli.Command{
				{
					Name: "add",
					Usage: "adds a new project",
					Action: AddProjectAction,
				},
			},
		},
	}

	app.Run(os.Args)
}