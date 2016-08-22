package main

import (
	"time"
	"math/rand"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	InitCLI()
}
