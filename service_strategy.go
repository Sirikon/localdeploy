package main

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
	return SystemdServiceStrategy{}
}
