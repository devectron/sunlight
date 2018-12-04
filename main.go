package main

import (
	"os"

	"github.com/devectron/sunlight/core"
	"github.com/devectron/sunlight/log"
)

func main() {
	log.Inf("Starting")
	core.StartListening(config())
}
func config() core.Config {
	serverPort := os.Getenv("SERVER_PORT")
	sqlPort := os.Getenv("SQL_PORT")
	sqlDbName := os.Getenv("SQL_DB_NAME")
	return core.Config{
		ServerPort: serverPort,
		SqlDbPort:  sqlPort,
		SqlDbName:  sqlDbName,
	}
}
