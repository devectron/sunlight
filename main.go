package main

import (
	"flag"
	"io"
	"net/http"
	"os"

	"github.com/devectron/sunlight/core"
	"github.com/devectron/sunlight/log"
)

var dbg bool

func init() {
	flag.BoolVar(&dbg, "dbg", false, "Debug")
	flag.Parse()
}

func main() {
	log.Inf("Starting [ Sunlight -v%s ]", core.VERSION)
	log.Dbg(dbg, "Debuging enabled")
	http.HandleFunc("/", index)
	err := http.ListenAndServe(":7375", nil)
	if err != nil {
		log.Err("%v", err)
	}
	//core.StartListening(config())
}
func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "<h1>No file with that name!</h1>")
}
func config() core.Config {
	serverPort := os.Getenv("SERVER_PORT")
	sqlPort := os.Getenv("SQL_PORT")
	sqlDbName := os.Getenv("SQL_DB_NAME")
	mailApiPublic := os.Getenv("MAILJET_PUBLIC")
	mailapiPrivate := os.Getenv("MAILJET_PRIVATE")
	convertApi := os.Getenv("CONVERT_API")
	if serverPort == "" {
		log.War("No $SERVER_PORT found using the default :7375")
		serverPort = "7375"
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
