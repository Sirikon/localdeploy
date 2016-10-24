package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os/exec"
	"fmt"
	"os"
	"io"
	"github.com/urfave/cli"
	"math/rand"
	"golang.org/x/crypto/bcrypt"
	"errors"
)

const Workspace = "/srv/molly"

type Project struct {
	Name string
	Config ProjectConfig
}

type ProjectConfig struct {
	Name string
	Token string
	Service string
}

func GetProjectByName(name string, project *Project) error {
	project.Name = name
	if err := project.LoadConfig(); err != nil {
		return errors.New("Specified project doesn't exist")
	}
	return nil
}

func (p *Project) CreateFilesFolder() error {
	filesFolder := Workspace + "/" + p.Name + "/files"

	if err := os.MkdirAll(filesFolder, 0777); err != nil {
		return err
	}

	return nil
}

func (p *Project) CleanFilesFolder() error {
	filesFolder := Workspace + "/" + p.Name + "/files"

	if err := os.RemoveAll(filesFolder); err != nil {
		return err
	}

	return p.CreateFilesFolder()
}

func (p *Project) GetEnvVars() []string {
	return []string{
		"MOLLY_ARTIFACT=" + Workspace + "/" + p.Name + "/artifact.zip",
	}
}

func (p *Project) RunDeploymentScript() error {
	if err := p.CleanFilesFolder(); err != nil {
		return errors.New("Couldn't clean the files folder: " + err.Error())
	}

	cmd := exec.Command("sh", Workspace + "/" + p.Name + "/deploy.sh")
	cmd.Dir = Workspace + "/" + p.Name + "/files/"
	cmd.Env = p.GetEnvVars()

	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error running the deployment script: " + err.Error() + "\n" + string(out))
	}

	return nil
}

func (p *Project) CreateDeploymentScript() error {
	return ioutil.WriteFile(Workspace + "/" + p.Name + "/deploy.sh", []byte("unzip $MOLLY_ARTIFACT\n"), 0700)
}

func (p *Project) CreateRunScript() error {
	return ioutil.WriteFile(Workspace + "/" + p.Name + "/run.sh", []byte("# Write here the run command\n"), 0700)
}

func (p *Project) StoreArtifact(fileReader io.Reader) error {
	var artifactFile *os.File

	if out, err := os.Create(Workspace + "/" + p.Name + "/artifact.zip"); err != nil {
		return errors.New("Couldn't create the artifact file: "  + err.Error())
	} else {
		artifactFile = out
	}

	defer artifactFile.Close()

	if _, err := io.Copy(artifactFile, fileReader); err != nil {
		return errors.New("Couldn't copy the artifact file to destiny: "  + err.Error())
	}

	return nil
}

func (p *Project) SaveConfig() error {
	projectFilePath := Workspace + "/" + p.Name + "/project.yml"

	var projectFileBytes []byte

	if out, err := yaml.Marshal(p.Config); err != nil {
		return err
	} else {
		projectFileBytes = out
	}

	if err := ioutil.WriteFile(projectFilePath, projectFileBytes, 0600); err != nil {
		return err
	}

	return nil
}

func (p *Project) LoadConfig() error {
	projectFilePath := Workspace + "/" + p.Name + "/project.yml"

	var projectFileBytes []byte

	if bytes, err := ioutil.ReadFile(projectFilePath); err != nil {
		return err
	} else {
		projectFileBytes = bytes
	}

	config := ProjectConfig{}

	if err := yaml.Unmarshal(projectFileBytes, &config); err != nil {
		return err
	}

	p.Config = config

	return nil
}

func (p Project) GetService() Service {
	return Service{Project:p}
}

func (p *Project) CreateService() error {
	return p.GetService().Save()
}

func (p *Project) RestartService() error {
	return p.GetService().Restart()
}

func GenerateRandomPassword() string {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, 32)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (p *Project) CheckToken(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.Config.Token), []byte(password))
	if err != nil {
		return false
	}
	return true
}

func AddProjectAction(c *cli.Context) error {
	projectName := c.Args().First()

	if projectName == "" {
		fmt.Println("Missing project name\n\nUsage:\nmolly project add [project name]")
		return nil
	}

	randomPassword := GenerateRandomPassword()
	hashedPassword, err := HashPassword(randomPassword)

	if err != nil {
		return err
	}

	project := Project{
		Name: projectName,
		Config: ProjectConfig{
			Name: projectName,
			Token: hashedPassword,
			Service: "molly-" + projectName,
		},
	}
	if err := project.CreateFilesFolder(); err != nil {
		return err
	}
	if err := project.CreateDeploymentScript(); err != nil {
		return err
	}
	if err := project.CreateRunScript(); err != nil {
		return err
	}
	if err := project.CreateService(); err != nil {
		fmt.Println("Error while creating the service")
		fmt.Println(err)
		return err
	}
	if err := project.SaveConfig(); err != nil {
		return err
	}

	fmt.Println("Automatically generated token:", randomPassword)

	return nil
}

func StartProjectServiceAction(c *cli.Context) error {
	projectName := c.Args().First()
	project := Project{}
	GetProjectByName(projectName, &project)
	project.RestartService()
	return nil
}