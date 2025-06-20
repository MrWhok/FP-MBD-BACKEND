package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func SendEmail(toEmail, subject, body string) error {
	from := os.Getenv("EMAIL_SENDER")       // ex: youraddress@gmail.com
	password := os.Getenv("EMAIL_PASSWORD") // App password dari Gmail

	// SMTP server configuration.
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	// Message.
	message := []byte(fmt.Sprintf("Subject: %s\r\n\r\n%s", subject, body))

	// Auth.
	auth := smtp.PlainAuth("", from, password, smtpHost)

	// Send email.
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, []string{toEmail}, message)
	return err
}
