package lib

import (
	"github.com/resend/resend-go/v2"
)

func Mail(from string, to []string, subject string, html string) error {
	client := resend.NewClient(Env.RESEND_KEY)

	params := &resend.SendEmailRequest{
		From:    from,
		To:      to,
		Subject: subject,
		Html:    html,
	}

	if _, err := client.Emails.Send(params); err != nil {
		return err
	}

	return nil
}
