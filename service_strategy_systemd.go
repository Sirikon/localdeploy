package main

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"os/user"
)

// SystemdServiceStrategy stablishes the logic needed to run Systemd
// services with molly
type SystemdServiceStrategy struct {
	ProjectPaths ProjectPaths
}

// Save the new service
func (p SystemdServiceStrategy) Save(service Service) error {
	currentUser, _ := user.Current()
	return ioutil.WriteFile("/etc/systemd/system/"+service.Project.Service+".service", []byte(`[Service]
WorkingDirectory=`+(p.ProjectPaths.GetFilesFolderPath(service.Project))+`
ExecStart=/bin/sh `+(p.ProjectPaths.GetRunScriptPath(service.Project))+`
Restart=always
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=`+(service.Project.Service)+`
User=`+(currentUser.Username)+`
Group=`+(currentUser.Username)+`
LimitNOFILE=64000

[Install]
WantedBy=multi-user.target
`), 0644)
}

// Start the service
func (p SystemdServiceStrategy) Start(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Project.Service + " start"
	cmd := exec.Command("/usr/sbin/service", service.Project.Service, "start")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to start the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}

// Stop the service
func (p SystemdServiceStrategy) Stop(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Project.Service + " stop"
	cmd := exec.Command("/usr/sbin/service", service.Project.Service, "stop")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to stop the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}

// Restart the service
func (p SystemdServiceStrategy) Restart(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Project.Service + " restart"
	cmd := exec.Command("/usr/sbin/service", service.Project.Service, "restart")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to restart the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}
