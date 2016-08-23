package main

import (
	"mime/multipart"
	"github.com/urfave/cli"
	"github.com/gin-gonic/gin"
	"errors"
)

func ActionDaemon(c *cli.Context) error {
	router := gin.Default()

	router.POST("/deploy", func(req *gin.Context) {

		project := &Project{}
		var uploadedFile multipart.File

		p := Promise{}
		p.Then(func () error {
			return GetProjectByName(req.PostForm("project"), project)
		})
		p.Then(func() error {
			if !project.CheckToken(req.PostForm("token")) {
				return errors.New("Wrong token")
			}
			return nil
		})
		p.Then(func() error {
			file, _, err := req.Request.FormFile("artifact")
			if err != nil {
				return errors.New("Error while reading the uploaded file")
			}
			uploadedFile = file
			return nil
		})
		p.Then(func() error {
			return project.StoreArtifact(uploadedFile)
		})
		p.Then(func() error {
			return project.RunDeploymentScript()
		})
		p.Then(func() error {
			return project.RestartService()
		})
		p.Then(func() error {
			req.String(200, "Done\n")
			return nil
		})
		p.Catch(func(err error) {
			req.String(400, err.Error() + "\n")
		})
		p.Run()

	})

	router.Run(":8080")

	return nil
}
