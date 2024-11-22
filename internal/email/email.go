package email

import (
	"fmt"
	"mailGOing/config"
	"net/smtp"
)

func SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", config.AppConfig.SMTPUser, config.AppConfig.SMTPPassword, config.AppConfig.SMTPHost)

	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\n\r\n%s",
		config.AppConfig.SMTPUser, to, subject, body)

	err := smtp.SendMail(
		fmt.Sprintf("%s:%d", config.AppConfig.SMTPHost, config.AppConfig.SMTPPort),
		auth,
		config.AppConfig.SMTPUser,
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		return fmt.Errorf("error sending email: %v", err)
	}
	return nil
}
