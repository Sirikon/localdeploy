package main

import (
	"github.com/urfave/cli"
	"os"
	"fmt"
)

func InitCLI() {

	cli.AppHelpTemplate = fmt.Sprintf(`
	   __     __)
  (, /|  /|      /) /)
    / | / |  ___// //
 ) /  |/  |_(_)(/_(/_ (_/_
(_/   '              .-/
                    (_/

%s`, cli.AppHelpTemplate)

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