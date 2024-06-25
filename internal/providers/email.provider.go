package providers

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mailersend/mailersend-go"
	"github.com/spf13/viper"
)

type EmailProviderInterface interface {
	SendWelcomeEmail(to string)
	SendForgotPasswordEmail(to string, validator uuid.UUID)
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

func (e *emailProvider) SendWelcomeEmail(to string) {
	if viper.Get("ENVIRONMENT") == "local" {
		fmt.Printf("Tried to send local welcome email to %s", to)
		return
	}

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
	message.SetSubject("Bem vindo ao na hora!")
	message.SetTemplateID(viper.Get("EMAIL_WELCOME_TEMPLATE_ID").(string))
	message.SetPersonalization([]mailersend.Personalization{
		{
			Email: to,
			Data: map[string]interface{}{
				"dashboard_url": fmt.Sprintf("%s/admin/login", viper.Get("WEB_URL")),
				"support_email": "contato@na-hora.com",
			},
		},
	})
	message.SetTags(tags)

	res, emailErr := e.sender.Email.Send(ctx, message)
	if emailErr != nil {
		fmt.Println(emailErr.Error())
	}

	fmt.Println(res)
}

func (e *emailProvider) SendForgotPasswordEmail(to string, validator uuid.UUID) {
	if viper.Get("ENVIRONMENT") == "local" {
		fmt.Printf("\nTried to send local forgot password email to %s\n\n", to)
		return
	}

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
	message.SetSubject("Recuperação de senha")
	message.SetTemplateID(viper.Get("EMAIL_FORGOT_PASSWORD_TEMPLATE_ID").(string))
	message.SetPersonalization([]mailersend.Personalization{
		{
			Email: to,
			Data: map[string]interface{}{
				"reset_link":    fmt.Sprintf("%s/admin/reset-password?email=%s&validator=%s", viper.Get("WEB_URL"), to, validator),
				"support_email": "contato@na-hora.com",
			},
		},
	})
	message.SetTags(tags)

	res, emailErr := e.sender.Email.Send(ctx, message)
	if emailErr != nil {
		fmt.Println(emailErr.Error())
	}

	fmt.Println(res)
}
