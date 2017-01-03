package main

import "os/exec"

// ICMD .
type ICMD interface {
	Exec(ExecParams) (string, error)
}

// CMD .
type CMD struct {
}

// ExecParams .
type ExecParams struct {
	Command []string
	CWD     string
	Env     []string
}

// Exec .
func (c *CMD) Exec(p ExecParams) (string, error) {
	cmd := exec.Command(p.Command[0], p.Command[1:]...)
	cmd.Dir = p.CWD
	cmd.Env = p.Env
	out, err := cmd.CombinedOutput()
	return string(out), err
}
