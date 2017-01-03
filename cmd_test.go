package main

import "github.com/stretchr/testify/mock"

type CMDMock struct {
	mock.Mock
}

func (m *CMDMock) Exec(a ExecParams) (string, error) {
	args := m.Called(a)
	return args.String(0), args.Error(1)
}
