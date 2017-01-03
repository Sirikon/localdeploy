package main

import yaml "gopkg.in/yaml.v2"

// IProjectSerialization .
type IProjectSerialization interface {
	Serialize(Project) ([]byte, error)
	Deserialize([]byte, *Project) error
}

// ProjectSerialization .
type ProjectSerialization struct {
}

// Serialize .
func (ps *ProjectSerialization) Serialize(project Project) ([]byte, error) {
	return yaml.Marshal(project)
}

// Deserialize .
func (ps *ProjectSerialization) Deserialize(bytes []byte, project *Project) error {
	return yaml.Unmarshal(bytes, project)
}
