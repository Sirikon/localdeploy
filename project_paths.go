package main

// IProjectPaths .
type IProjectPaths interface {
	GetHomePath(Project) string
	GetFilePath(Project) string
	GetFilesFolderPath(Project) string
	GetDeploymentScriptPath(Project) string
	GetRunScriptPath(Project) string
	GetArtifactPath(Project) string
}

// ProjectPaths path builder for projects
type ProjectPaths struct {
	Config Config
}

// GetHomePath returns the project's home path
func (p ProjectPaths) GetHomePath(project Project) string {
	return p.Config.Workspace + p.Config.PathSep + project.Name
}

// GetFilePath returns the project file definition path
func (p ProjectPaths) GetFilePath(project Project) string {
	return p.GetHomePath(project) + p.Config.PathSep + "project.yml"
}

// GetFilesFolderPath returns the files folder path
func (p ProjectPaths) GetFilesFolderPath(project Project) string {
	return p.GetHomePath(project) + p.Config.PathSep + "files"
}

// GetDeploymentScriptPath returns the deploy script path
func (p ProjectPaths) GetDeploymentScriptPath(project Project) string {
	return p.GetHomePath(project) + p.Config.PathSep + "deploy.sh"
}

// GetRunScriptPath returns the run script path
func (p ProjectPaths) GetRunScriptPath(project Project) string {
	return p.GetHomePath(project) + p.Config.PathSep + "run.sh"
}

// GetArtifactPath returns the artifact path
func (p ProjectPaths) GetArtifactPath(project Project) string {
	return p.GetHomePath(project) + p.Config.PathSep + "artifact.zip"
}
