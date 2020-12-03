package main

import (
	"gopkg.in/gomail.v2"
	"log"
)

func main() {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", "danastudio@datatom.com", "")         // 发件人
	m.SetHeader("To", m.FormatAddress("hongbiaoc289@gmail.com", "")) // 收件人
	m.SetHeader("Subject", "google邮箱")                               // 主题
	m.SetBody("text/html", "向google邮箱发送消息")
	d := gomail.NewDialer("smtp.exmail.qq.com", 465, "danastudio@datatom.com", "DTnanjing666") // 发送邮件服务器、端口、发件人账号、发件人密码
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("邮件发送失败", err)
	}
}
