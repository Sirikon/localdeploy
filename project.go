package main

// Workspace is the default folder where all the
// molly projects are stored
const Workspace = "/srv/molly"

// Project defines a project
type Project struct {
	Name    string
	Token   string
	Service string
}

// GetHomePath returns the project's home path
func (p Project) GetHomePath() string {
	return Workspace + "/" + p.Name
}

// GetFilePath returns the project file definition path
func (p Project) GetFilePath() string {
	return p.GetHomePath() + "/project.yml"
}

// GetFilesFolderPath returns the files folder path
func (p Project) GetFilesFolderPath() string {
	return p.GetHomePath() + "/files"
}

// GetDeploymentScriptPath returns the deploy script path
func (p Project) GetDeploymentScriptPath() string {
	return p.GetHomePath() + "/deploy.sh"
}

// GetRunScriptPath returns the run script path
func (p Project) GetRunScriptPath() string {
	return p.GetHomePath() + "/run.sh"
}

// GetArtifactPath returns the artifact path
func (p Project) GetArtifactPath() string {
	return p.GetHomePath() + "/artifact.zip"
}

// GetDeploymentEnvVars returns the environment vars needed
// for deploy script
func (p Project) GetDeploymentEnvVars() []string {
	return []string{
		"MOLLY_ARTIFACT=" + p.GetArtifactPath(),
	}
}
