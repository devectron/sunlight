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
	VERSION = "0.9.0"
	AUTHOR  = "hihebark"
	MAIL    = "n.amara@prtonmail.ch"
)
