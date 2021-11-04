package util

import (
	"Go-User-System/config"
	"crypto/tls"
	"fmt"
	"net/smtp"
	"strconv"
)

func SendEmail(to, subject, body string) (err error) {
	//连接邮箱
	conn, err := tls.Dial("tcp", config.Config.SMTP.Server+":"+strconv.Itoa(config.Config.SMTP.Port), nil)
	if err != nil {
		return
	}
	mailConn, err := smtp.NewClient(conn, config.Config.SMTP.Server)
	if err != nil {
		return
	}

	//拼接内容
	header := make(map[string]string)
	header["From"] = config.Config.SMTP.Address
	header["To"] = to
	header["Subject"] = subject
	header["Content-Type"] = "text/html;charset=UTF-8"
	message := ""
	for key, value := range header {
		message += fmt.Sprintf("%s:%s\r\n", key, value)
	}
	message += "\r\n" + body

	//发送
	auth := smtp.PlainAuth(
		"",
		config.Config.SMTP.Address,
		config.Config.SMTP.Password,
		config.Config.SMTP.Server,
	)
	if err = mailConn.Auth(auth); err != nil {
		return
	}
	if err = mailConn.Mail(config.Config.SMTP.Address); err != nil {
		return
	}
	if err = mailConn.Rcpt(to); err != nil {
		fmt.Print(err)
		return
	}
	w, err := mailConn.Data()
	if err != nil {
		return
	}
	_, err = w.Write([]byte(message))
	if err != nil {
		return
	}
	err = w.Close()
	if err != nil {
		return
	}
	return mailConn.Quit()
}
