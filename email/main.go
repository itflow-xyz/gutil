package email

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"
)

var (
	smtpServer   = os.Getenv("SMTP_SERVER")
	smtpPort     = os.Getenv("SMTP_PORT")
	smtpUsername = os.Getenv("SMTP_USERNAME")
	smtpPassword = os.Getenv("SMTP_PASSWORD")
)

func SendEmail(dest, object, content string) error {
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

func SendEmailTSL(dest, subject, body string) error {
	from := mail.Address{Name: "", Address: smtpUsername}
	to := mail.Address{Name: "", Address: dest}

	// Setup headers
	headers := make(map[string]string)
	headers["From"] = from.String()
	headers["To"] = to.String()
	headers["Subject"] = subject

	// Setup message
	message := ""
	for k, v := range headers {
		message += fmt.Sprintf("%s: %s\r\n", k, v)
	}
	message += "\r\n" + body

	// Connect to the SMTP Server
	servername := smtpServer + ":" + smtpPort //"mail.infomaniak.com:465"

	auth := smtp.PlainAuth("", smtpUsername, smtpPassword, smtpServer)

	// TLS config
	tlsconfig := &tls.Config{
		InsecureSkipVerify: true,
		ServerName:         smtpServer,
	}

	var err error
	// Here is the key, you need to call tls.Dial instead of smtp.Dial
	// for smtp servers running on 465 that require an ssl connection
	// from the very beginning (no starttls)
	conn, err := tls.Dial("tcp", servername, tlsconfig)
	if err != nil {
		log.Panic(err)
	}

	c, err := smtp.NewClient(conn, smtpServer)
	if err != nil {
		log.Panic(err)
	}

	// Auth
	if err = c.Auth(auth); err != nil {
		log.Panic(err)
	}

	// To && From
	if err = c.Mail(from.Address); err != nil {
		log.Panic(err)
	}

	if err = c.Rcpt(to.Address); err != nil {
		log.Panic(err)
	}

	// Data
	w, err := c.Data()
	if err != nil {
		log.Panic(err)
	}

	_, err = w.Write([]byte(message))
	if err != nil {
		log.Panic(err)
	}

	err = w.Close()
	if err != nil {
		log.Panic(err)
	}

	c.Quit()

	return err
}
