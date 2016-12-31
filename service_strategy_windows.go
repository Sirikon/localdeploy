package main

// WindowsServiceStrategy stablishes the logic needed to run Windows
// services with molly
type WindowsServiceStrategy struct {
}

// Save the new service
func (p WindowsServiceStrategy) Save(service Service) error {
	return nil
}

// Start the service
func (p WindowsServiceStrategy) Start(service Service) error {
	return nil
}

// Stop the service
func (p WindowsServiceStrategy) Stop(service Service) error {
	return nil
}

// Restart the service
func (p WindowsServiceStrategy) Restart(service Service) error {
	return nil
}
