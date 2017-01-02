package main

import (
	"fmt"

	"errors"

	"github.com/urfave/cli"
)

// IProjectActions .
type IProjectActions interface {
	AddCLIAction(*cli.Context) error
	AddAction(string) error
	StartServiceCLIAction(*cli.Context) error
	StartServiceAction(string) error
}

// ProjectActions .
type ProjectActions struct {
	projectLogic IProjectLogic
}

func (pa *ProjectActions) validateAddContext(c *cli.Context) bool {
	projectName := c.Args().First()
	if projectName == "" {
		fmt.Println("Missing project name\n\nUsage:\nmolly project add [project name]")
		return false
	}
	return true
}

// AddCLIAction .
func (pa *ProjectActions) AddCLIAction(c *cli.Context) error {
	if !pa.validateAddContext(c) {
		return nil
	}
	return pa.AddAction(c.Args().First())
}

// AddAction will create a new project in system
func (pa *ProjectActions) AddAction(projectName string) error {

	if pa.projectLogic.Exists(projectName) {
		return errors.New("Project " + projectName + " already exists")
	}

	randomToken := pa.projectLogic.GenerateRandomToken()
	var hashedToken string
	if hash, err := pa.projectLogic.HashToken(randomToken); err == nil {
		hashedToken = hash
	} else {
		return err
	}

	project := Project{
		Name:    projectName,
		Token:   hashedToken,
		Service: "molly-" + projectName,
	}

	if err := pa.projectLogic.CreateFilesFolder(project); err != nil {
		return errors.New("Error while creating the files folder:\n" + err.Error())
	}
	if err := pa.projectLogic.CreateDeploymentScript(project); err != nil {
		return errors.New("Error while creating the deployment script:\n" + err.Error())
	}
	if err := pa.projectLogic.CreateRunScript(project); err != nil {
		return errors.New("Error while creating the run script:\n" + err.Error())
	}
	if err := pa.projectLogic.CreateService(project); err != nil {
		return errors.New("Error while creating the service:\n" + err.Error())
	}
	if err := pa.projectLogic.Save(project); err != nil {
		return errors.New("Error while creating the project.yml file:\n" + err.Error())
	}

	fmt.Println("Automatically generated token:", randomToken)

	return nil
}

func (pa *ProjectActions) validateStartServiceContext(c *cli.Context) bool {
	projectName := c.Args().First()
	if projectName == "" {
		fmt.Println("Missing project name\n\nUsage:\nmolly project service start [project name]")
		return false
	}
	return true
}

// StartServiceCLIAction .
func (pa *ProjectActions) StartServiceCLIAction(c *cli.Context) error {
	if !pa.validateStartServiceContext(c) {
		return nil
	}
	return pa.StartServiceAction(c.Args().First())
}

// StartServiceAction starts the project's service
func (pa *ProjectActions) StartServiceAction(projectName string) error {
	project := Project{}

	if err := pa.projectLogic.GetByName(projectName, &project); err != nil {
		return err
	}

	if err := pa.projectLogic.RestartService(project); err != nil {
		return err
	}

	return nil
}
