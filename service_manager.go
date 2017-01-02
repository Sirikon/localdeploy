package main

import (
	"runtime"
)

// IServiceManager interface
type IServiceManager interface {
	Save(Project) error
	Start(Project) error
	Stop(Project) error
	Restart(Project) error
}

// InitServiceManager .
func InitServiceManager() IServiceManager {
	if runtime.GOOS == "windows" {
		return WindowsServiceManager{}
	}
	return SystemdServiceManager{}
}
