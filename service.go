package main

type Service struct {
	Name string
}

func (s Service) Save() error {
	return GetServiceStrategy().Save(s)
}

func (s Service) Start() error {
	return GetServiceStrategy().Start(s)
}

func (s Service) Stop() error {
	return GetServiceStrategy().Stop(s)
}

func (s Service) Restart() error {
	return GetServiceStrategy().Restart(s)
}
