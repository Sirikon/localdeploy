package main

import (
	"os/user"
	"io/ioutil"
	"os/exec"
	"errors"
	"fmt"
)

type SystemdServiceStrategy struct {

}

func (p SystemdServiceStrategy) Save(service Service) error {
	currentUser, _ := user.Current()
	return ioutil.WriteFile("/etc/systemd/system/" + service.Name + ".service", []byte(`[Service]
WorkingDirectory=` + (Workspace + "/" + service.Name + "/files") + `
ExecStart=/bin/sh ` + (Workspace + "/" + service.Name + "/run.sh") + `
Restart=always
StandardOutput=syslog
StandardError=syslog
SyslogIdentifier=` + ("molly-" + service.Name) + `
User=` + (currentUser.Username) + `
Group=` + (currentUser.Username) + `
LimitNOFILE=64000

[Install]
WantedBy=multi-user.target
`), 0644)
}

func (p SystemdServiceStrategy) Start(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Name + " start"
	fmt.Println("Command to execute: /usr/sbin/service " + service.Name + " start")
	cmd := exec.Command("/usr/sbin/service", service.Name, "start")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to restart the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}

func (p SystemdServiceStrategy) Stop(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Name + " stop"
	fmt.Println("Command to execute: /usr/sbin/service " + service.Name + " stop")
	cmd := exec.Command("/usr/sbin/service", service.Name, "stop")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to restart the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}

func (p SystemdServiceStrategy) Restart(service Service) error {
	commandToExecute := "Command to execute: /usr/sbin/service " + service.Name + " restart"
	cmd := exec.Command("/usr/sbin/service", service.Name, "restart")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return errors.New("There was an error while trying to restart the service: " + err.Error() + "\n" + string(out) + "\n\nUsed command:\n" + commandToExecute)
	}
	return nil
}
