package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"io"
	"io/ioutil"
	"gopkg.in/yaml.v2"
	"os/exec"
	"fmt"
)

type ProjectConfig struct {
	Name string
	Token string
	Service string
}

func getProjectConfig(projectName string) ProjectConfig {
	bytes, _ := ioutil.ReadFile("/srv/localdeploy/projects/" + projectName + "/project.yml")

	config := ProjectConfig{}

	yaml.Unmarshal(bytes, &config)

	return config
}

func getProjectToken(projectName string) string {
	return getProjectConfig(projectName).Token
}

func getProjectServiceName(projectName string) string {
	return getProjectConfig(projectName).Service
}

func cleanProjectFilesFolder(projectName string) error {
	os.RemoveAll("/srv/localdeploy/projects/" + projectName + "/files")
	err := os.Mkdir("/srv/localdeploy/projects/" + projectName + "/files", 0777)
	if err != nil {
		fmt.Println("Error cleaning project")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func runProjectDeploymentScript(projectName string) error {
	cleanErr := cleanProjectFilesFolder(projectName)
	if cleanErr != nil {
		return cleanErr
	}
	cmd := exec.Command("sh", "/srv/localdeploy/projects/" + projectName + "/deploy.sh")
	cmd.Dir = "/srv/localdeploy/projects/" + projectName + "/files/"
	cmd.Env = []string{"LD_ARTIFACT=/srv/localdeploy/artifacts/" + projectName + "/artifact.zip"}
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error running deployment script")
		fmt.Println(string(out))
		return err
	}
	return nil
}

func restartProject(projectName string) error {
	cmd := exec.Command("/usr/sbin/service", getProjectServiceName(projectName), "restart")
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error restarting service")
		fmt.Println(string(out))
		return err
	}
	return nil
}

func storeProjectArtifact(projectName string, fileReader io.Reader) error {
	out, err := os.Create("/srv/localdeploy/artifacts/" + projectName + "/artifact.zip")
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, fileReader)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	router := gin.Default()

	router.POST("/upload/:project", func(c *gin.Context) {

		project := c.Param("project")
		token := c.Query("token")
		validProjectToken := getProjectToken(project)

		if token != validProjectToken || validProjectToken == "" {
			c.String(400, "Nope")
			return
		}

		file, _, err := c.Request.FormFile("upload")
		if err != nil {
			c.String(500, "Error")
			return
		}

		err2 := storeProjectArtifact(project, file)
		if err2 != nil {
			c.String(500, "Error")
			return
		}

		err3 := runProjectDeploymentScript(project)
		if err3 != nil {
			c.String(500, "Error running project deploy script")
			return
		}

		err4 := restartProject(project)
		if err4 != nil {
			c.String(500, "Error restarting service")
			return
		}

		c.String(200, "Done")
	})
	router.Run(":8080")
}