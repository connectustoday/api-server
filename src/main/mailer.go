package main

import (
	"bytes"
	"crypto/tls"
	"errors"
	"html/template"
	"log"
	"net/smtp"
	"strconv"
)

// Send mail using template

func SendMail(recipient string, subject string, fromTemplate string, replace interface{}) error {
	t := template.Must(template.New("_all").Parse(fromTemplate))

	var body bytes.Buffer
	if err := t.Execute(&body, replace); err != nil {
		return err
	}

	msg := []byte("To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n\r\n" +
		body.String() + "\r\n")

	//println(string(msg))
	return smtp.SendMail(SMTP_HOST+":"+strconv.Itoa(SMTP_PORT), loginAuthSMTP(MAIL_USERNAME, MAIL_PASSWORD), MAIL_SENDER, []string{recipient}, msg)
}

// Test the SMTP connection

func InitMailer() {
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SMTP_HOST,
	}

	Mailer, err := smtp.Dial(SMTP_HOST + ":" + strconv.Itoa(SMTP_PORT))
	if err != nil {
		log.Fatal(err)
	}

	if err = Mailer.StartTLS(tlsconfig); err != nil {
		log.Fatal(err)
	}

	if err = Mailer.Auth(loginAuthSMTP(MAIL_USERNAME, MAIL_PASSWORD)); err != nil {
		log.Fatal(err)
	}

	if err = Mailer.Mail(MAIL_SENDER); err != nil {
		log.Fatal(err)
	}
	log.Println("Verified SMTP configuration!")
}

// External utility for using AUTH LOGIN instead of PLAIN AUTH
// https://gist.github.com/andelf/5118732

type loginAuth struct {
	username, password string
}

func loginAuthSMTP(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("Unknown from server")
		}
	}
	return nil, nil
}
