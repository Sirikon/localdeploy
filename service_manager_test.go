package main

import (
	"github.com/stretchr/testify/mock"
)

type ServiceManagerMock struct {
	mock.Mock
}

func (m *ServiceManagerMock) Save(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ServiceManagerMock) Start(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ServiceManagerMock) Stop(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
func (m *ServiceManagerMock) Restart(a Project) error {
	args := m.Called(a)
	return args.Error(0)
}
