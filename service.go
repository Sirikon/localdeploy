package main

// Service defines a project's Service
// it's mainly a passthru which uses the
// proper implementation depending on the OS
type Service struct {
	Project Project
}

// Save .
func (s Service) Save() error {
	return GetServiceStrategy().Save(s)
}

// Start .
func (s Service) Start() error {
	return GetServiceStrategy().Start(s)
}

// Stop .
func (s Service) Stop() error {
	return GetServiceStrategy().Stop(s)
}

// Restart .
func (s Service) Restart() error {
	return GetServiceStrategy().Restart(s)
}
