package mail

import (
	"gopkg.in/gomail.v2"
	"github.com/vimsucks/resmtp/model"
)

func SendMail(msg *model.Message, dl *model.Dialer) (err error) {
	if msg.ContentType == "" {
		msg.ContentType = "text/plain"
	}
	if msg.From == "" {
		msg.From = dl.Username
	}
	if msg.FromName == "" {
		msg.FromName = dl.FromName
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(msg.From, msg.FromName))
	m.SetHeader("To", msg.To...)
	m.SetHeader("Subject", msg.Subject)
	m.SetBody(msg.ContentType, msg.Content)
	for _, attach := range msg.Attach {
		m.Attach(attach)
	}

	d := gomail.NewPlainDialer(dl.Host, dl.Port, dl.Username, dl.Password)
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
