package main

import (
	"os"

	"github.com/stretchr/testify/mock"
)

type FileSystemMock struct {
	mock.Mock
}

func (m *FileSystemMock) WriteFile(a string, b []byte, c os.FileMode) error {
	args := m.Called(a, b, c)
	return args.Error(0)
}
func (m *FileSystemMock) ReadFile(a string) ([]byte, error) {
	args := m.Called(a)
	return args.Get(0).([]byte), args.Error(1)
}
func (m *FileSystemMock) MkdirAll(a string, b os.FileMode) error {
	args := m.Called(a, b)
	return args.Error(0)
}
func (m *FileSystemMock) RemoveAll(a string) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *FileSystemMock) Create(a string) (*os.File, error) {
	args := m.Called(a)
	return args.Get(0).(*os.File), args.Error(1)
}
