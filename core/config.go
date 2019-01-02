package core

//Config configuration port.
type Config struct {
	ServerPort     string
	SqlDbPort      string
	SqlDbName      string
	EmailName      string
	MailApiPublic  string
	MailApiPrivate string
	ConvertApi     string
	DBG            bool
}

const (
	VERSION = "1.0.0"
	AUTHOR  = "[hihebark] and the teams"
	MAIL    = "n.amara@prtonmail.ch"
)
