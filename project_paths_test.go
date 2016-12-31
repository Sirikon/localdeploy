package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProjectPaths(t *testing.T) {
	project := Project{Name: "test"}
	pp := ProjectPaths{Config{PathSep: "/", Workspace: "/srv/molly"}}

	assert.Equal(t, pp.GetFilePath(project), "/srv/molly/test/project.yml")
	assert.Equal(t, pp.GetFilesFolderPath(project), "/srv/molly/test/files")
	assert.Equal(t, pp.GetDeploymentScriptPath(project), "/srv/molly/test/deploy.sh")
	assert.Equal(t, pp.GetRunScriptPath(project), "/srv/molly/test/run.sh")
	assert.Equal(t, pp.GetArtifactPath(project), "/srv/molly/test/artifact.zip")
}
