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

	m, err := getMailer()
	if err != nil {
		return err
	}

	defer m.Quit()

	msg := []byte("From: " + MAIL_SENDER +"\r\n" +
		"To: " + recipient + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n\r\n" +
		body.String() + "\r\n")

	if err := m.Rcpt(recipient); err != nil {
		return err
	}

	w, err := m.Data()
	if err != nil {
		return err
	}

	_, err = w.Write(msg)
	if err != nil {
		return err
	}

	_ = w.Close()
	if err != nil {
		return err
	}

	return nil
}

func getMailer() (*smtp.Client, error) {
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         SMTP_HOST,
	}

	m, err := smtp.Dial(SMTP_HOST + ":" + strconv.Itoa(SMTP_PORT))
	if err != nil {
		return nil, err
	}

	if SMTP_TLS {
		if err = m.StartTLS(tlsconfig); err != nil {
			return nil, err
		}

		if err = m.Auth(loginAuthSMTP(MAIL_USERNAME, MAIL_PASSWORD)); err != nil {
			return nil, err
		}
	} else {
		if err = m.Auth(smtp.PlainAuth("", MAIL_USERNAME, MAIL_PASSWORD, SMTP_HOST)); err != nil {
			return nil, err
		}
	}

	if err = m.Mail(MAIL_SENDER); err != nil {
		return nil, err
	}
	return m, nil
}

// Test the SMTP connection

func InitMailer(startup bool) {
	if _, err := getMailer(); err != nil {
		log.Fatal(err)
	}

	if startup {
		log.Println("Verified SMTP configuration!")
	}
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
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}
