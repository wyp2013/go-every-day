package model

import (
	"fmt"
	"github.com/mailgun/mailgun-go"
	"context"
	"os"
	"time"
)



type Mailgun struct {
	domain string
	apiKey string
	mailgun mailgun.Mailgun
}

func NewMaigunModel(domain, apiKey string) *Mailgun {
	mg := mailgun.NewMailgun(domain, apiKey)

	return &Mailgun{
		domain:domain,
		apiKey:apiKey,
		mailgun:mg,
	}
}


func (m *Mailgun) SendMessage(from, subject, text string, to []string, filePathList []string) (string, error) {
	message := m.mailgun.NewMessage(from, subject, text, to...)

	for _, file := range filePathList {
		_, err := os.Stat(file)
		if err != nil {
			fmt.Println(fmt.Sprintf("file:%s not exist!", file))
			continue
		}

		message.AddAttachment(file)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*30)
	defer cancel()

	msg, id, err := m.mailgun.Send(ctx, message)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	fmt.Println(msg, id)
	return id, nil
}
