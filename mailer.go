package api_server

import (
	"crypto/tls"
	"log"
	"net/smtp"
	"strconv"
)

func InitMailer() {
	auth := smtp.PlainAuth("", MAIL_USERNAME, MAIL_PASSWORD, SMTP_HOST)
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: SMTP_HOST,
	}

	c, err := tls.Dial("tcp", SMTP_HOST + ":" + strconv.Itoa(SMTP_PORT), tlsconfig)
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	Mailer, err = smtp.NewClient(c, SMTP_HOST)
	if err != nil {
		log.Fatal(err)
	}
	defer Mailer.Close()

	if err = Mailer.Auth(auth); err != nil {
		log.Fatal(err)
	}

	if err = Mailer.Mail(MAIL_SENDER); err != nil {
		
	}

}