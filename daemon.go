package main

import (
	"errors"
	"mime/multipart"

	"bytes"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli"
)

// IDaemonAction .
type IDaemonAction interface {
	Run(*cli.Context) error
}

// DaemonAction .
type DaemonAction struct {
	projectLogic IProjectLogic
}

// Run defines a CLI action which initializes the daemon
// which listens HTTP requests
func (da DaemonAction) Run(c *cli.Context) error {

	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()

	router.POST("/deploy", func(req *gin.Context) {

		project := Project{}
		var uploadedFile multipart.File
		var deploymentLog bytes.Buffer
		projectName := req.PostForm("project")
		receivedToken := req.PostForm("token")

		deploymentLog.WriteString("=== Starting Deployment ===\n")

		p := Promise{}
		p.Then(func() error {
			deploymentLog.WriteString("- Getting project with name '" + projectName + "'\n")
			if projectName == "" {
				return errors.New("Project name can't be empty")
			}
			if err := da.projectLogic.GetByName(projectName, &project); err != nil {
				return errors.New("Project '" + projectName + "' doesn't exist")
			}
			return da.projectLogic.GetByName(projectName, &project)
		})
		p.Then(func() error {
			deploymentLog.WriteString("- Validating received token\n")
			var tokenIsCorrect = da.projectLogic.CheckToken(project, receivedToken)
			if !tokenIsCorrect {
				return errors.New("Wrong token")
			}
			return nil
		})
		p.Then(func() error {
			deploymentLog.WriteString("- Reading received artifact\n")
			file, _, err := req.Request.FormFile("artifact")
			if err != nil {
				return errors.New("Error while reading the uploaded file")
			}
			uploadedFile = file
			return nil
		})
		p.Then(func() error {
			deploymentLog.WriteString("- Storing received artifact\n")
			return da.projectLogic.StoreArtifact(project, uploadedFile)
		})
		p.Then(func() error {
			deploymentLog.WriteString("- Running deployment script\n")
			return da.projectLogic.RunDeploymentScript(project)
		})
		p.Then(func() error {
			deploymentLog.WriteString("- Restarting the service\n")
			return da.projectLogic.RestartService(project)
		})
		p.Then(func() error {
			deploymentLog.WriteString("Done!\n")
			req.String(200, deploymentLog.String())
			return nil
		})
		p.Catch(func(err error) {
			deploymentLog.WriteString("There was an ERROR:\n\n")
			deploymentLog.WriteString(err.Error() + "\n")
			req.String(400, deploymentLog.String())
		})
		p.Run()

	})

	router.Run(":8080")

	return nil
}
