package email

import (
	"fmt"
	"net/smtp"
	"os"
)

var (
	smtpServer   = os.Getenv("SMTP_SERVER")
	smtpPort     = os.Getenv("SMTP_PORT")
	smtpUsername = os.Getenv("SMTP_USERNAME")
	smtpPassword = os.Getenv("SMTP_PASSWORD")
)

func sendEmail(dest, object, content string) error {
	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)

	to := []string{dest}
	email := buildEmail(dest, smtpUsername, object, content)
	return smtp.SendMail(smtpServer+":"+smtpPort, auth, dest, to, email)
}

func buildEmail(to, from, object, content string) []byte {
	email := "To: %s\r\n"
	email += "From: %s\r\n"
	email += "Subject: %s\r\n"
	email += "\r\n"
	email += "%s\r\n"

	return []byte(fmt.Sprintf(to, from, object, content))
}
