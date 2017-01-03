package main

import "github.com/stretchr/testify/mock"

type ProjectSerializationMock struct {
	mock.Mock
}

func (m *ProjectSerializationMock) Serialize(a Project) ([]byte, error) {
	args := m.Called(a)
	return args.Get(0).([]byte), args.Error(1)
}

// Deserialize .
func (m *ProjectSerializationMock) Deserialize(a []byte, b *Project) error {
	args := m.Called(a, b)
	return args.Error(0)
}
