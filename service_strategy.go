package main

type ServiceStrategy interface {
	Save(Service) error
	Start(Service) error
	Stop(Service) error
	Restart(Service) error
}

func GetServiceStrategy() ServiceStrategy {
	return SystemdServiceStrategy{}
}
