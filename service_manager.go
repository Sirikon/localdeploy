package main

// IServiceManager interface
type IServiceManager interface {
	Save(Project) error
	Start(Project) error
	Stop(Project) error
	Restart(Project) error
}
