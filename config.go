package main

import (
	"os"
	"runtime"
)

// Config .
type Config struct {
	OS        string
	PathSep   string
	Workspace string
}

// Init initializes the configuration based on
// runtime info and environment variables
func (c *Config) Init() {
	c.OS = runtime.GOOS
	c.PathSep = string(os.PathSeparator)
	var workspace = os.Getenv("MOLLY_WORKSPACE")
	if workspace != "" {
		c.Workspace = workspace
	} else {
		c.Workspace = "/srv/molly"
	}
}
