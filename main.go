package main

import (
	"flag"
	"os"

	"github.com/devectron/sunlight/core"
	"github.com/devectron/sunlight/log"
	"google.golang.org/appengine"
)

var dbg bool

func init() {
	flag.BoolVar(&dbg, "dbg", false, "Debug")
	flag.Parse()
	if _, err := os.Stat("tmp"); os.IsNotExist(err) {
		os.Mkdir("tmp", 0666)
	}
}

func main() {
	log.Inf("Starting [ Sunlight -v%s ]", core.VERSION)
	log.Dbg(dbg, "Debuging enabled")
	core.StartListening(config())
	appengine.Main()
}
func config() core.Config {
	serverPort := os.Getenv("SERVER_PORT")
	sqlPort := os.Getenv("PORT")
	sqlDbName := os.Getenv("SQL_DB_NAME")
	mailApiPublic := os.Getenv("MAILJET_PUBLIC")
	mailapiPrivate := os.Getenv("MAILJET_PRIVATE")
	convertApi := os.Getenv("CONVERT_API")
	if serverPort == "" {
		log.War("No $SERVER_PORT found using the default :5000")
		serverPort = "5000"
	}
	if sqlPort == "" {
		log.War("No $SQL_PORT found.")
	}
	if sqlDbName == "" {
		log.War("No $SQL_DB_NAME found.")
	}
	if mailApiPublic == "" {
		log.War("No $MAILJET_PUBLIC found.")
	}
	if mailapiPrivate == "" {
		log.War("No $MAILJET_PRIVATE found.")
	}
	if convertApi == "" {
		log.War("No $CONVERT_API found.")
	}
	return core.Config{
		ServerPort:     serverPort,
		SqlDbPort:      sqlPort,
		SqlDbName:      sqlDbName,
		EmailName:      "devectron.not.replay@gmail.com",
		MailApiPublic:  mailApiPublic,
		MailApiPrivate: mailapiPrivate,
		ConvertApi:     convertApi,
		DBG:            dbg,
	}
}
