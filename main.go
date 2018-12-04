package main

import (
	"github.com/devectron/sunlight/core"
	"github.com/devectron/sunlight/log"
)

func main() {
	log.Inf("Starting")
	core.StartListening()
}
