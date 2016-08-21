package main

import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os/exec"
	"fmt"
	"os"
	"io"
	"github.com/urfave/cli"
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

func CreateProject(name string) error {

}

func GetProjectByName(name string, project *Project) error {
	project.Name = name
	if err := project.LoadConfig(); err != nil {
		return err
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
		return err
	}

	cmd := exec.Command("sh", Workspace + "/" + p.Name + "/deploy.sh")
	cmd.Dir = Workspace + "/" + p.Name + "/files/"
	cmd.Env = p.GetEnvVars()

	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}

	return nil
}

func (p *Project) StoreArtifact(fileReader io.Reader) error {
	var artifactFile *os.File

	if out, err := os.Create(Workspace + "/" + p.Name + "/artifact.zip"); err != nil {
		return err
	} else {
		artifactFile = out
	}

	defer artifactFile.Close()

	if _, err := io.Copy(artifactFile, fileReader); err != nil {
		return err
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

	if err := ioutil.WriteFile(projectFilePath, projectFileBytes, 0777); err != nil {
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

func (p *Project) RestartService() error {
	cmd := exec.Command("/usr/sbin/service", p.Config.Service, "restart")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(string(out))
		return err
	}
	return nil
}

func AddProjectAction(c *cli.Context) error {

	projectName := c.Args().First()

	project := Project{
		Name: projectName,
		Config: ProjectConfig{
			Name: projectName,
			Token: "lalala",
			Service: projectName,
		},
	}
	if err := project.CreateFilesFolder(); err != nil {
		return err
	}
	if err := project.SaveConfig(); err != nil {
		return err
	}

	return nil
}