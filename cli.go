package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
)

var appVersion string

// CLI .
type CLI struct {
	daemonAction   IDaemonAction
	projectActions IProjectActions
}

// Init initializes the CLI options
func (c CLI) Init() {

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
	app.Version = appVersion

	app.Commands = []cli.Command{
		{
			Name:   "daemon",
			Usage:  "Starts the application as a daemon",
			Action: c.daemonAction.Run,
		},
		{
			Name:  "project",
			Usage: "Project management",
			Subcommands: []cli.Command{
				{
					Name:   "add",
					Usage:  "adds a new project",
					Action: c.projectActions.AddCLIAction,
				},
				{
					Name:  "service",
					Usage: "Project service management",
					Subcommands: []cli.Command{
						{
							Name:   "start",
							Usage:  "starts the service",
							Action: c.projectActions.StartServiceCLIAction,
						},
					},
				},
			},
		},
	}

	app.Run(os.Args)
}
