package main

import (
	"errors"
	"io"
	"math/rand"
	"os"
	"os/exec"

	"golang.org/x/crypto/bcrypt"

	yaml "gopkg.in/yaml.v2"
)

// IProjectLogic .
type IProjectLogic interface {
	GetByName(string, *Project) error
	CreateFilesFolder(Project) error
	CleanFilesFolder(Project) error
	RunDeploymentScript(Project) error
	CreateDeploymentScript(Project) error
	CreateRunScript(Project) error
	StoreArtifact(Project, io.Reader) error
	GenerateRandomToken() string
	HashToken(string) (string, error)
	CreateService(Project) error
	Save(Project) error
	RestartService(Project) error
	CheckToken(Project, string) bool
}

// ProjectLogic logic related to projects
type ProjectLogic struct {
	Config         Config
	projectPaths   IProjectPaths
	serviceManager IServiceManager
	fileSystem     IFileSystem
}

// GetByName populates the given project with the existing one
func (pl ProjectLogic) GetByName(projectName string, project *Project) error {

	var projectFileBytes []byte

	if bytes, err := pl.readProjectFile(projectName); err == nil {
		projectFileBytes = bytes
	} else {
		return err
	}

	if err := yaml.Unmarshal(projectFileBytes, &project); err != nil {
		return err
	}

	return nil
}

// CreateFilesFolder creates a 'files' folder inside the
// project's folder
func (pl ProjectLogic) CreateFilesFolder(project Project) error {
	var filesFolderPath = pl.projectPaths.GetFilesFolderPath(project)
	if err := pl.fileSystem.MkdirAll(filesFolderPath, 0777); err != nil {
		return err
	}
	return nil
}

// CleanFilesFolder will clean all the content inside the
// project's files folder
func (pl ProjectLogic) CleanFilesFolder(project Project) error {
	var filesFolderPath = pl.projectPaths.GetFilesFolderPath(project)
	if err := pl.fileSystem.RemoveAll(filesFolderPath); err != nil {
		return err
	}
	return pl.CreateFilesFolder(project)
}

// RunDeploymentScript will execute the project's deployment script (deploy.sh)
func (pl ProjectLogic) RunDeploymentScript(project Project) error {

	if err := pl.CleanFilesFolder(project); err != nil {
		return errors.New("Couldn't clean the files folder: " + err.Error())
	}

	var deploymentScriptPath = pl.projectPaths.GetDeploymentScriptPath(project)
	var filesFolderPath = pl.projectPaths.GetFilesFolderPath(project)

	cmd := exec.Command("sh", deploymentScriptPath)
	cmd.Dir = filesFolderPath
	cmd.Env = pl.getDeploymentEnvVars(project)

	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error running the deployment script: " + err.Error() + "\n" + string(out))
	}

	return nil
}

// CreateDeploymentScript created the default deploy script
func (pl ProjectLogic) CreateDeploymentScript(project Project) error {
	var deploymentScriptContent = "unzip $MOLLY_ARTIFACT\n"
	var deploymentScriptPath = pl.projectPaths.GetDeploymentScriptPath(project)
	return pl.fileSystem.WriteFile(deploymentScriptPath, []byte(deploymentScriptContent), 0700)
}

// CreateRunScript creates the default run script
func (pl ProjectLogic) CreateRunScript(project Project) error {
	var runScriptContent = "# Write here the run command\n"
	var runScriptPath = pl.projectPaths.GetRunScriptPath(project)
	return pl.fileSystem.WriteFile(runScriptPath, []byte(runScriptContent), 0700)
}

// StoreArtifact will store the new artifact
func (pl ProjectLogic) StoreArtifact(project Project, fileReader io.Reader) error {
	var artifactFile *os.File
	var artifactPath = pl.projectPaths.GetArtifactPath(project)

	if out, err := pl.fileSystem.Create(artifactPath); err == nil {
		artifactFile = out
	} else {
		return errors.New("Couldn't create the artifact file: " + err.Error())
	}

	defer artifactFile.Close()

	if _, err := pl.fileSystem.Copy(artifactFile, fileReader); err != nil {
		return errors.New("Couldn't copy the artifact file to destiny: " + err.Error())
	}

	return nil
}

// GenerateRandomToken will generate a random token
func (pl ProjectLogic) GenerateRandomToken() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// HashToken hashes a token
func (pl ProjectLogic) HashToken(token string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedToken), nil
}

// CreateService creates the project's service in host system
func (pl ProjectLogic) CreateService(project Project) error {
	return pl.serviceManager.Save(project)
}

// Save writes to disk the project file
func (pl ProjectLogic) Save(project Project) error {

	var projectFileBytes []byte

	if out, err := yaml.Marshal(project); err == nil {
		projectFileBytes = out
	} else {
		return err
	}

	var projectFilePath = pl.projectPaths.GetFilePath(project)

	if err := pl.fileSystem.WriteFile(projectFilePath, projectFileBytes, 0600); err != nil {
		return err
	}

	return nil
}

// RestartService restarts the project's service
func (pl ProjectLogic) RestartService(project Project) error {
	return pl.serviceManager.Restart(project)
}

// CheckToken checks if a token is valid for the project
func (pl ProjectLogic) CheckToken(project Project, token string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(project.Token), []byte(token))
	if err != nil {
		return false
	}
	return true
}

func (pl ProjectLogic) getDeploymentEnvVars(project Project) []string {
	return []string{
		"MOLLY_ARTIFACT=" + pl.projectPaths.GetArtifactPath(project),
	}
}

func (pl ProjectLogic) readProjectFile(projectName string) ([]byte, error) {
	var projectFilePath = pl.projectPaths.GetFilePath(Project{Name: projectName})
	return pl.fileSystem.ReadFile(projectFilePath)
}
