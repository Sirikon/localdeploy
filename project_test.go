package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectGetFilePath(t *testing.T) {
	project := Project{Name: "test"}

	var filePath = project.GetFilePath()
	var expectedFilePath = "/srv/molly/test/project.yml"

	assert.Equal(t, filePath, expectedFilePath)
}

func TestProjectGetDeployEnvVars(t *testing.T) {
	project := Project{Name: "test"}

	var envVars = project.GetDeploymentEnvVars()

	assert.Equal(t, len(envVars), 1)
	assert.Equal(t, envVars[0], "MOLLY_ARTIFACT="+project.GetArtifactPath())
}
