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

// InitConfig initializes the configuration based on
// runtime info and environment variables
func InitConfig() Config {
	c := Config{}
	c.OS = runtime.GOOS
	c.PathSep = string(os.PathSeparator)
	c.Workspace = initConfigWorkspace()
	return c
}

func initConfigWorkspace() string {
	var workspace = os.Getenv("MOLLY_WORKSPACE")
	if workspace != "" {
		return workspace
	}

	if runtime.GOOS == "windows" {
		return "C:\\molly"
	}
	return "/srv/molly"
}
