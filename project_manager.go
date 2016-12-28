package main

import (
	"errors"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"

	"golang.org/x/crypto/bcrypt"

	yaml "gopkg.in/yaml.v2"
)

// GetProjectByName populates the given project with the existing one
func GetProjectByName(projectName string, project *Project) error {

	var projectFileBytes []byte

	if bytes, err := readProjectFile(projectName); err == nil {
		projectFileBytes = bytes
	} else {
		return err
	}

	if err := yaml.Unmarshal(projectFileBytes, &project); err != nil {
		return err
	}

	return nil
}

// CreateProjectFilesFolder creates a 'files' folder inside the
// project's folder
func CreateProjectFilesFolder(project Project) error {
	if err := os.MkdirAll(project.GetFilesFolderPath(), 0777); err != nil {
		return err
	}
	return nil
}

// CleanProjectFilesFolder will clean all the content inside the
// project's files folder
func CleanProjectFilesFolder(project Project) error {
	if err := os.RemoveAll(project.GetFilesFolderPath()); err != nil {
		return err
	}
	return CreateProjectFilesFolder(project)
}

// RunProjectDeploymentScript will execute the project's deployment script (deploy.sh)
func RunProjectDeploymentScript(project Project) error {

	if err := CleanProjectFilesFolder(project); err != nil {
		return errors.New("Couldn't clean the files folder: " + err.Error())
	}

	cmd := exec.Command("sh", project.GetDeploymentScriptPath())
	cmd.Dir = project.GetFilesFolderPath()
	cmd.Env = project.GetDeploymentEnvVars()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error running the deployment script: " + err.Error() + "\n" + string(out))
	}

	return nil
}

// CreateProjectDeploymentScript created the default deploy script
func CreateProjectDeploymentScript(project Project) error {
	var deploymentScriptContent = "unzip $MOLLY_ARTIFACT\n"
	return ioutil.WriteFile(project.GetDeploymentScriptPath(), []byte(deploymentScriptContent), 0700)
}

// CreateProjectRunScript creates the default run script
func CreateProjectRunScript(project Project) error {
	var runScriptContent = "# Write here the run command\n"
	return ioutil.WriteFile(project.GetRunScriptPath(), []byte(runScriptContent), 0700)
}

// StoreProjectArtifact will store the new artifact
func StoreProjectArtifact(project Project, fileReader io.Reader) error {
	var artifactFile *os.File

	if out, err := os.Create(project.GetArtifactPath()); err == nil {
		artifactFile = out
	} else {
		return errors.New("Couldn't create the artifact file: " + err.Error())
	}

	defer artifactFile.Close()

	if _, err := io.Copy(artifactFile, fileReader); err != nil {
		return errors.New("Couldn't copy the artifact file to destiny: " + err.Error())
	}

	return nil
}

// GenerateRandomToken will generate a random token
func GenerateRandomToken() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

// HashToken hashes a token
func HashToken(token string) (string, error) {
	hashedToken, err := bcrypt.GenerateFromPassword([]byte(token), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedToken), nil
}

// CreateProjectService creates the project's service in host system
func CreateProjectService(project Project) error {
	return getProjectService(project).Save()
}

// WriteProjectFile writes to disk the project file
func WriteProjectFile(project Project) error {

	var projectFileBytes []byte

	if out, err := yaml.Marshal(project); err == nil {
		projectFileBytes = out
	} else {
		return err
	}

	if err := ioutil.WriteFile(project.GetFilePath(), projectFileBytes, 0600); err != nil {
		return err
	}

	return nil
}

// RestartProjectService restarts the project's service
func RestartProjectService(project Project) error {
	return getProjectService(project).Restart()
}

// CheckProjectToken checks if a token is valid for the project
func CheckProjectToken(project Project, token string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(project.Token), []byte(token))
	if err != nil {
		return false
	}
	return true
}

func getProjectService(project Project) Service {
	return Service{Project: project}
}

func readProjectFile(projectName string) ([]byte, error) {
	return ioutil.ReadFile(Project{Name: projectName}.GetFilePath())
}
