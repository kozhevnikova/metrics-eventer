package main

import (
	"fmt"
	"os"

	gomail "gopkg.in/gomail.v2"
)

func sendMail(config Config, message string, email string) {
	m := gomail.NewMessage()
	m.SetHeader("From", config.Mail.AddressFrom)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Metrics eventer")
	m.SetBody("", message)

	d := gomail.NewDialer(
		config.Mail.ServerName,
		config.Mail.Port,
		config.Mail.AddressFrom,
		config.Mail.Password,
	)

	if err := d.DialAndSend(m); err != nil {
		fmt.Fprintln(os.Stderr, "(SEND MAIL) ERROR:", err)
		return
	}
}
