package main

import (
	"io"

	"github.com/stretchr/testify/mock"
)

type ProjectLogicMock struct {
	mock.Mock
}

func (m *ProjectLogicMock) GetByName(a string, b *Project) error {
	args := m.Called(a, b)
	return args.Error(0)
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
