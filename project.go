package main

import (
	"os"
)

// Workspace is the default folder where all the
// molly projects are stored
var Workspace = os.Getenv("MOLLY_WORKSPACE")

//const Workspace = "/srv/molly"

// Project defines a project
type Project struct {
	Name    string
	Token   string
	Service string
}
