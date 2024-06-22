package providers

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
	"github.com/spf13/viper"
)

type EmailProviderInterface interface {
	SendWelcomeEmail(to string, subject string, html string)
	SendForgotPasswordEmail(to string, subject string, html string)
}

type emailProvider struct {
	sender *mailersend.Mailersend
}

func NewEmailProvider() EmailProviderInterface {
	var APIKey = viper.Get("MAIL_SENDER_TOKEN")

	if APIKey == nil {
		panic("MAIL_SENDER_TOKEN not set")
	}

	APITokenStringified := fmt.Sprintf("%v", APIKey)
	sender := mailersend.NewMailersend(APITokenStringified)

	return &emailProvider{
		sender,
	}
}

func (e *emailProvider) SendWelcomeEmail(to string, subject string, html string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  "Na hora",
		Email: "admin@na-hora.com",
	}

	recipients := []mailersend.Recipient{
		{
			Email: to,
		},
	}

	tags := []string{"welcome"}

	message := e.sender.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetTags(tags)

	res, emailErr := e.sender.Email.Send(ctx, message)
	if emailErr != nil {
		fmt.Println(emailErr.Error())
	}

	fmt.Println(res)
}

func (e *emailProvider) SendForgotPasswordEmail(to string, subject string, html string) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	from := mailersend.From{
		Name:  "Na hora",
		Email: "admin@na-hora.com",
	}

	recipients := []mailersend.Recipient{
		{
			Email: to,
		},
	}

	tags := []string{"forgot-password"}

	message := e.sender.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetTags(tags)

	res, emailErr := e.sender.Email.Send(ctx, message)
	if emailErr != nil {
		fmt.Println(emailErr.Error())
	}

	fmt.Println(res)
}
