package main

import (
	"os/user"
	"io/ioutil"
	"os/exec"
	"errors"
)

type SystemdServiceStrategy struct {

}

func (p SystemdServiceStrategy) Save(service Service) error {
	currentUser, _ := user.Current()
	return ioutil.WriteFile("/etc/systemd/system/" + service.Project.Config.Service + ".service", []byte(`[Service]
WorkingDirectory=` + (Workspace + "/" + service.Project.Config.Name + "/files") + `
ExecStart=/bin/sh ` + (Workspace + "/" + service.Project.Config.Name + "/run.sh") + `
Restart=always
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=` + (service.Project.Config.Service) + `
User=` + (currentUser.Username) + `
Group=` + (currentUser.Username) + `
LimitNOFILE=64000

[Install]
WantedBy=multi-user.target
`), 0644)
}

func (p SystemdServiceStrategy) Start(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Project.Config.Service + " start"
	cmd := exec.Command("/usr/sbin/service", service.Project.Config.Service, "start")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to start the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}

func (p SystemdServiceStrategy) Stop(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Project.Config.Service + " stop"
	cmd := exec.Command("/usr/sbin/service", service.Project.Config.Service, "stop")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to stop the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}

func (p SystemdServiceStrategy) Restart(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Project.Config.Service + " restart"
	cmd := exec.Command("/usr/sbin/service", service.Project.Config.Service, "restart")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to restart the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}
