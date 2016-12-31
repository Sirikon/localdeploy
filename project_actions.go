package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// ProjectActions .
type ProjectActions struct {
	ProjectLogic ProjectLogic
}

func (pa ProjectActions) validateAddContext(c *cli.Context) bool {
	projectName := c.Args().First()
	if projectName == "" {
		fmt.Println("Missing project name\n\nUsage:\nmolly project add [project name]")
		return false
	}
	return true
}

// AddAction will create a new project in system
func (pa ProjectActions) AddAction(c *cli.Context) error {
	if !pa.validateAddContext(c) {
		return nil
	}

	projectName := c.Args().First()

	randomToken := pa.ProjectLogic.GenerateRandomToken()
	hashedToken, err := pa.ProjectLogic.HashToken(randomToken)

	if err != nil {
		return err
	}

	project := Project{
		Name:    projectName,
		Token:   hashedToken,
		Service: "molly-" + projectName,
	}
	if err := pa.ProjectLogic.CreateFilesFolder(project); err != nil {
		return err
	}
	if err := pa.ProjectLogic.CreateDeploymentScript(project); err != nil {
		return err
	}
	if err := pa.ProjectLogic.CreateRunScript(project); err != nil {
		return err
	}
	if err := pa.ProjectLogic.CreateService(project); err != nil {
		fmt.Println("Error while creating the service")
		fmt.Println(err)
		return err
	}
	if err := pa.ProjectLogic.Save(project); err != nil {
		return err
	}

	fmt.Println("Automatically generated token:", randomToken)

	return nil
}

func (pa ProjectActions) validateStartServiceContext(c *cli.Context) bool {
	projectName := c.Args().First()
	if projectName == "" {
		fmt.Println("Missing project name\n\nUsage:\nmolly project service start [project name]")
		return false
	}
	return true
}

// StartServiceAction starts the project's service
func (pa ProjectActions) StartServiceAction(c *cli.Context) error {
	if !pa.validateStartServiceContext(c) {
		return nil
	}

	projectName := c.Args().First()

	project := Project{}
	pa.ProjectLogic.GetByName(projectName, &project)
	pa.ProjectLogic.RestartService(project)
	return nil
}
