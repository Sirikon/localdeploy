package main

import (
	"runtime"
)

// ServiceStrategy interface
type ServiceStrategy interface {
	Save(Service) error
	Start(Service) error
	Stop(Service) error
	Restart(Service) error
}

// GetServiceStrategy returns the proper ServiceStrategy
// instance depending on the host system
func GetServiceStrategy() ServiceStrategy {
	if runtime.GOOS == "windows" {
		return WindowsServiceStrategy{}
	}
	return SystemdServiceStrategy{}
}
