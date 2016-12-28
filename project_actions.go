package main

import (
	"fmt"

	"github.com/urfave/cli"
)

func validateAddProjectContext(c *cli.Context) bool {
	projectName := c.Args().First()
	if projectName == "" {
		fmt.Println("Missing project name\n\nUsage:\nmolly project add [project name]")
		return false
	}
	return true
}

// AddProjectAction will create a new project in system
func AddProjectAction(c *cli.Context) error {
	if !validateAddProjectContext(c) {
		return nil
	}

	projectName := c.Args().First()

	randomToken := GenerateRandomToken()
	hashedToken, err := HashToken(randomToken)

	if err != nil {
		return err
	}

	project := Project{
		Name:    projectName,
		Token:   hashedToken,
		Service: "molly-" + projectName,
	}
	if err := CreateProjectFilesFolder(project); err != nil {
		return err
	}
	if err := CreateProjectDeploymentScript(project); err != nil {
		return err
	}
	if err := CreateProjectRunScript(project); err != nil {
		return err
	}
	if err := CreateProjectService(project); err != nil {
		fmt.Println("Error while creating the service")
		fmt.Println(err)
		return err
	}
	if err := WriteProjectFile(project); err != nil {
		return err
	}

	fmt.Println("Automatically generated token:", randomToken)

	return nil
}

func validateStartProjectServiceContext(c *cli.Context) bool {
	projectName := c.Args().First()
	if projectName == "" {
		fmt.Println("Missing project name\n\nUsage:\nmolly project service start [project name]")
		return false
	}
	return true
}

// StartProjectServiceAction starts the project's service
func StartProjectServiceAction(c *cli.Context) error {
	if !validateStartProjectServiceContext(c) {
		return nil
	}

	projectName := c.Args().First()

	project := Project{}
	GetProjectByName(projectName, &project)
	RestartProjectService(project)
	return nil
}
