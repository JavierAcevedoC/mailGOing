package email

import (
	"bytes"
	"fmt"
	"html/template"
	"mailGOing/config"
	"net/smtp"
)

type EmailData struct {
	To      string
	Subject string
	Body    string
}

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

func SendEmailWithTemplate(to, subject, body string) error {
	tmpl, err := template.ParseFiles("internal/email/templates/default.html")
	if err != nil {
		return fmt.Errorf("Error parsing default template: %v", err)
	}
	data := EmailData{
		To:      to,
		Subject: subject,
		Body:    body,
	}
	var render bytes.Buffer
	if err := tmpl.Execute(&render, data); err != nil {
		return fmt.Errorf("Error rendering default template: %v", err)
	}

	auth := smtp.PlainAuth("", config.AppConfig.SMTPUser, config.AppConfig.SMTPPassword, config.AppConfig.SMTPHost)
	msg := fmt.Sprintf("From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/html; charset=UTF-8\r\n\r\n%s",
		config.AppConfig.SMTPUser, to, subject, render.String())
	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", config.AppConfig.SMTPHost, config.AppConfig.SMTPPort),
		auth,
		config.AppConfig.SMTPUser,
		[]string{to},
		[]byte(msg),
	)

	if err != nil {
		return fmt.Errorf("Error sending email with default template: %v", err)
	}

	return nil
}
