package main

import (
	"crypto/tls"
	"fmt"

	"gopkg.in/gomail.v2"
)

type EmailConfig struct {
	ToAddress      string  `json:"to_address"`
	Subject        string  `json:"subject"`
	ServerAddress  string  `json:"server_address"`
	ServerPort     float64 `json:"server_port"`
	ServerUsername string  `json:"server_username"`
	ServerPassword string  `json:"server_password"`
}

func emailPublish(message string, config map[string]interface{}) error {
	emailFooterMessager := `
		<br><br>
		This is an automated email sent via Flightpath
	`

	var eConf EmailConfig
	eConf.ToAddress = config["to_address"].(string)
	eConf.Subject = config["subject"].(string)
	eConf.ServerAddress = config["server_address"].(string)
	eConf.ServerPort = config["server_port"].(float64)
	eConf.ServerUsername = config["server_username"].(string)
	eConf.ServerPassword = config["server_password"].(string)

	d := gomail.NewDialer(eConf.ServerAddress, int(eConf.ServerPort), eConf.ServerUsername, eConf.ServerPassword)
	d.TLSConfig = &tls.Config{}
	d.TLSConfig.InsecureSkipVerify = true

	m := gomail.NewMessage()
	m.SetHeader("From", eConf.ServerUsername)
	m.SetHeader("To", eConf.ToAddress)
	m.SetHeader("Subject", eConf.Subject)
	m.SetBody("text/html", fmt.Sprint(message, emailFooterMessager))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	e(fmt.Sprint("Email sent to ", eConf.ToAddress, " with subject ", eConf.Subject))

	return nil
}

func newEmailOutlet(id string, listenFrom []string) *OutletModule {
	return &OutletModule{
		Name:       "email",
		ID:         id,
		ListenFrom: listenFrom,
		Publish:    emailPublish,
		configured: false,
	}
}

func registerEmailOutlet() {
	e := newDocsEntry("email", NodeTypeOutlet)
	e.AddParam("TO Address", "string", "Who to send the email to", "to_address")
	e.AddParam("Subject", "string", "The email's subject", "subject")
	e.AddParam("Server Address", "string", "SMTP Server address", "server_address")
	e.AddParam("Server Port", "int", "SMTP Server port", "server_port")
	e.AddParam("Server Username", "string", "SMTP Server username, the email you want to send the message from", "server_username")
	e.AddParam("Server Password", "string", "SMTP Server password", "server_password")
	AppDocs = append(AppDocs, e)
}
