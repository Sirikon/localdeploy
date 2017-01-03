package main

import (
	"io"

	"testing"

	"os"

	"errors"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ProjectLogicMock struct {
	mock.Mock
}

func (m *ProjectLogicMock) GetByName(a string, b *Project) error {
	args := m.Called(a, b)
	return args.Error(0)
}
func (m *ProjectLogicMock) Exists(a string) bool {
	args := m.Called(a)
	return args.Bool(0)
}
func (m *ProjectLogicMock) CreateFilesFolder(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ProjectLogicMock) CleanFilesFolder(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ProjectLogicMock) RunDeploymentScript(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ProjectLogicMock) CreateDeploymentScript(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ProjectLogicMock) CreateRunScript(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ProjectLogicMock) StoreArtifact(a Project, b io.Reader) error {
	args := m.Called(a, b)
	return args.Error(0)
}
func (m *ProjectLogicMock) GenerateRandomToken() string {
	args := m.Called()
	return args.String(0)
}
func (m *ProjectLogicMock) HashToken(a string) (string, error) {
	args := m.Called(a)
	return args.String(0), args.Error(1)
}
func (m *ProjectLogicMock) CreateService(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ProjectLogicMock) Save(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ProjectLogicMock) RestartService(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ProjectLogicMock) CheckToken(a Project, b string) bool {
	args := m.Called(a, b)
	return args.Bool(0)
}

func getProjectLogicMockedDependencies() (Config, *ProjectPaths, *ProjectSerializationMock, *ServiceManagerMock, *FileSystemMock, *CMDMock) {
	var config = Config{}
	config.OS = "linux"
	config.PathSep = "/"
	config.Workspace = "/srv/molly"

	var projectPaths = &ProjectPaths{config}
	var projectSerialization = &ProjectSerializationMock{}
	var serviceManager = &ServiceManagerMock{}
	var fileSystem = &FileSystemMock{}
	var cmd = &CMDMock{}

	return config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd
}

func TestProjectLogic_Exists(t *testing.T) {
	config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd := getProjectLogicMockedDependencies()

	var project = Project{}
	var projectBytes = []byte{0xAA, 0xBB, 0xCC}

	// Should read the file from the correct place and call a deserialization
	fileSystem.On("ReadFile", "/srv/molly/test/project.yml").Once().Return(projectBytes, nil)
	projectSerialization.On("Deserialize", projectBytes, &project).Once().Return(nil)

	var projectLogic = ProjectLogic{config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd}

	projectLogic.Exists("test")

	fileSystem.AssertExpectations(t)
	projectSerialization.AssertExpectations(t)
}

func TestProjectLogic_CreateFilesFolderAndRunDeploymentScript(t *testing.T) {
	config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd := getProjectLogicMockedDependencies()

	var project = Project{
		Name: "test",
	}

	var folderPerm os.FileMode = 0777
	var execParams = ExecParams{
		Command: []string{"sh", "/srv/molly/test/deploy.sh"},
		CWD:     "/srv/molly/test/files",
		Env:     []string{"MOLLY_ARTIFACT=/srv/molly/test/artifact.zip"},
	}
	// Should create all directories in directory mode
	fileSystem.On("MkdirAll", "/srv/molly/test/files", folderPerm).Times(2).Return(nil)
	fileSystem.On("RemoveAll", "/srv/molly/test/files").Once().Return(nil)
	cmd.On("Exec", execParams).Once().Return("", nil)

	var projectLogic = ProjectLogic{config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd}

	projectLogic.CreateFilesFolder(project)
	projectLogic.RunDeploymentScript(project)

	fileSystem.AssertExpectations(t)
}

func TestProjectLogic_RunDeploymentScriptWithError(t *testing.T) {
	config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd := getProjectLogicMockedDependencies()

	var project = Project{
		Name: "test",
	}

	var folderPerm os.FileMode = 0777
	var execParams = ExecParams{
		Command: []string{"sh", "/srv/molly/test/deploy.sh"},
		CWD:     "/srv/molly/test/files",
		Env:     []string{"MOLLY_ARTIFACT=/srv/molly/test/artifact.zip"},
	}

	fileSystem.On("MkdirAll", "/srv/molly/test/files", folderPerm).Once().Return(nil)
	fileSystem.On("RemoveAll", "/srv/molly/test/files").Once().Return(nil)
	cmd.On("Exec", execParams).Once().Return("OUTPUT_ERROR", errors.New("EXECUTION_ERROR"))

	var projectLogic = ProjectLogic{config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd}

	err := projectLogic.RunDeploymentScript(project)

	assert.EqualError(t, err, "There was an error running the deployment script:\n\n EXECUTION_ERROR\n\nCommand Output:\nOUTPUT_ERROR")

	fileSystem.AssertExpectations(t)
}

func TestProjectLogic_CreateRunAndDeploymentScripts(t *testing.T) {
	config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd := getProjectLogicMockedDependencies()

	var project = Project{
		Name: "test",
	}

	var filePerm os.FileMode = 0700
	fileSystem.On("WriteFile", "/srv/molly/test/run.sh", []byte("# Write here the run command\n"), filePerm).Once().Return(nil)
	fileSystem.On("WriteFile", "/srv/molly/test/deploy.sh", []byte("unzip $MOLLY_ARTIFACT\n"), filePerm).Once().Return(nil)

	var projectLogic = ProjectLogic{config, projectPaths, projectSerialization, serviceManager, fileSystem, cmd}

	projectLogic.CreateRunScript(project)
	projectLogic.CreateDeploymentScript(project)

	fileSystem.AssertExpectations(t)
}
