package email

import (
	"strings"

	"gopkg.in/gomail.v2"
)

type Client interface {
	Send(r *Request) error
}

type Request struct {
	MailTo  string // 收件人 多个用,分割
	Subject string // 邮件主题
	Body    string // 邮件内容
}

var _ Client = (*EmailClient)(nil)

type Options struct {
	MailHost string
	MailPort int
	MailUser string // 发件人
	MailPass string // 发件人密码
}

type EmailClient struct {
	Options
	*gomail.Dialer
}

func NewEmailClient(opt Options) *EmailClient {
	client := &EmailClient{Options: opt}
	client.Dialer = gomail.NewDialer(client.MailHost, client.MailPort, client.MailUser, client.MailPass)
	return client
}

func (client EmailClient) Send(r *Request) error {

	m := gomail.NewMessage()

	//设置发件人
	m.SetHeader("From", client.MailUser)

	//设置发送给多个用户
	mailArrTo := strings.Split(r.MailTo, ",")
	m.SetHeader("To", mailArrTo...)

	//设置邮件主题
	m.SetHeader("Subject", r.Subject)

	//设置邮件正文
	m.SetBody("text/html", r.Body)

	return client.DialAndSend(m)
}
