package util

import (
	"Go-User-System/config"
	"net/smtp"
	"strconv"
)

func SendEmail(to, subject, body string) error {
	auth := smtp.PlainAuth("", config.Config.SMTP.Address, config.Config.SMTP.Password, config.Config.SMTP.Server)
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + config.Config.SMTP.Address + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + body)
	err := smtp.SendMail(config.Config.SMTP.Server+":"+strconv.Itoa(config.Config.SMTP.Port), auth, config.Config.SMTP.Address, []string{to}, msg)
	return err
}
