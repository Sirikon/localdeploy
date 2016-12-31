package main

import (
	"fmt"

	"github.com/urfave/cli"
)

// IProjectActions .
type IProjectActions interface {
	AddAction(*cli.Context) error
	StartServiceAction(*cli.Context) error
}

// ProjectActions .
type ProjectActions struct {
	projectLogic IProjectLogic
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

	randomToken := pa.projectLogic.GenerateRandomToken()
	hashedToken, err := pa.projectLogic.HashToken(randomToken)

	if err != nil {
		return err
	}

	project := Project{
		Name:    projectName,
		Token:   hashedToken,
		Service: "molly-" + projectName,
	}
	if err := pa.projectLogic.CreateFilesFolder(project); err != nil {
		return err
	}
	if err := pa.projectLogic.CreateDeploymentScript(project); err != nil {
		return err
	}
	if err := pa.projectLogic.CreateRunScript(project); err != nil {
		return err
	}
	if err := pa.projectLogic.CreateService(project); err != nil {
		fmt.Println("Error while creating the service")
		fmt.Println(err)
		return err
	}
	if err := pa.projectLogic.Save(project); err != nil {
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
	pa.projectLogic.GetByName(projectName, &project)
	pa.projectLogic.RestartService(project)
	return nil
}
