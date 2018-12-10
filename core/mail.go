package core

import (
	"fmt"
	"os"

	"github.com/devectron/sunlight/log"
	. "github.com/mailjet/mailjet-apiv3-go"
)

func SendMail(tomail string, content string) {
	log.Inf("Sending e-mail to %s", tomail)
	mailjetClient := NewMailjetClient(os.Getenv("MJ_APIKEY_PUBLIC"), os.Getenv("MJ_APIKEY_PRIVATE"))
	email := &InfoSendMail{
		FromEmail: "pilot@mailjet.com",
		FromName:  "Mailjet Pilot",
		Subject:   "Your email flight plan!",
		TextPart:  "Dear passenger, welcome to Mailjet! May the delivery force be with you!",
		HTMLPart:  "<h3>Dear passenger, welcome to Mailjet!</h3><br />May the delivery force be with you!",
		Recipients: []Recipient{
			Recipient{
				Email: "passenger@mailjet.com",
			},
		},
	}
	res, err := mailjetClient.SendMail(email)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success")
		fmt.Println(res)
	}
}
