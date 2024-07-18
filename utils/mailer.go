package mailer

import (
	"log"
	"os"
	"strconv"

	"github.com/go-mail/mail"
)

func SendMail(to, subject, body string) error {
	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("EMAIL"))
	m.SetHeader("To", to)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	port, err := strconv.Atoi(os.Getenv("SMTP_PORT"))
	if err != nil {
		log.Println("Invalid SMTP_PORT:", os.Getenv("SMTP_PORT"))
		return err
	}

	d := mail.NewDialer(os.Getenv("SMTP_SERVER"), port, os.Getenv("EMAIL"), os.Getenv("EMAIL_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
}
