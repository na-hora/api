package providers

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
	"github.com/spf13/viper"
)

type EmailProvider interface {
	SendWelcomeEmail(to string, subject string, html string)
}

type emailProvider struct{}

func NewEmailProvider() EmailProvider {
	return &emailProvider{}
}

func (e *emailProvider) SendWelcomeEmail(to string, subject string, html string) {
	var APIKey = viper.Get("MAIL_SENDER_TOKEN")

	if APIKey == nil {
		panic("MAIL_SENDER_TOKEN not set")
	}

	APITokenStringified := fmt.Sprintf("%v", APIKey)

	ms := mailersend.NewMailersend(APITokenStringified)

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

	message := ms.Email.NewMessage()

	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(html)
	message.SetTags(tags)

	res, emailErr := ms.Email.Send(ctx, message)
	if emailErr != nil {
		fmt.Println(emailErr.Error())
	}

	fmt.Println(res)
}
