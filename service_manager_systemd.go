package main

import (
	"errors"
	"io/ioutil"
	"os/exec"
	"os/user"
)

// SystemdServiceManager stablishes the logic needed to run Systemd
// services with molly
type SystemdServiceManager struct {
	ProjectPaths ProjectPaths
}

// Save the new service
func (p SystemdServiceManager) Save(project Project) error {
	currentUser, _ := user.Current()
	return ioutil.WriteFile("/etc/systemd/system/"+project.Service+".service", []byte(`[Service]
WorkingDirectory=`+(p.ProjectPaths.GetFilesFolderPath(project))+`
ExecStart=/bin/sh `+(p.ProjectPaths.GetRunScriptPath(project))+`
Restart=always
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=`+(project.Service)+`
User=`+(currentUser.Username)+`
Group=`+(currentUser.Username)+`
LimitNOFILE=64000

[Install]
WantedBy=multi-user.target
`), 0644)
}

// Start the service
func (p SystemdServiceManager) Start(project Project) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + project.Service + " start"
	cmd := exec.Command("/usr/sbin/service", project.Service, "start")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to start the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}

// Stop the service
func (p SystemdServiceManager) Stop(project Project) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + project.Service + " stop"
	cmd := exec.Command("/usr/sbin/service", project.Service, "stop")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to stop the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}

// Restart the service
func (p SystemdServiceManager) Restart(project Project) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + project.Service + " restart"
	cmd := exec.Command("/usr/sbin/service", project.Service, "restart")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to restart the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}
