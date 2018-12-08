package main

import (
	"bytes"
	"crypto/tls"
	"html/template"
	"log"
	"net/smtp"
	"strconv"
)

func SendMail(recipient string, subject string, fromTemplate string, replace interface{}) error {
	t := template.Must(template.New("temp").Parse(fromTemplate))

	var body bytes.Buffer
	if err := t.Execute(&body, replace); err != nil {
		return err
	}

	auth := smtp.PlainAuth("", MAIL_USERNAME, MAIL_PASSWORD, SMTP_HOST)
	to := []string{recipient}
	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"\r\n" +
		body.String() + "\r\n")

	return smtp.SendMail(SMTP_HOST + ":" + strconv.Itoa(SMTP_PORT), auth, MAIL_SENDER, to, msg)
}

func InitMailer() {
	auth := smtp.PlainAuth("", MAIL_USERNAME, MAIL_PASSWORD, SMTP_HOST)
	tlsconfig := &tls.Config {
		InsecureSkipVerify: true,
		ServerName: SMTP_HOST,
	}

	Mailer, err := smtp.Dial(SMTP_HOST + ":" + strconv.Itoa(SMTP_PORT))
	if err != nil {
		log.Fatal(err)
	}

	err = Mailer.StartTLS(tlsconfig)
	if err != nil {
		log.Fatal(err)
	}

	if err = Mailer.Auth(auth); err != nil {
		log.Fatal(err)
	}

	if err = Mailer.Mail(MAIL_SENDER); err != nil {
		log.Fatal(err)
	}
	log.Println("Verified SMTP configuration!")
}