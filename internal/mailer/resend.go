package mailer

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/resend/resend-go/v2"
)

type resendMailer struct {
	client *resend.Client
	sender EmailData
}

func NewResend(fromName, fromEmail, apiKey string) Client {
	client := resend.NewClient(apiKey)
	return &resendMailer{
		client: client,
		sender: EmailData{
			Name:  fromName,
			Email: fromEmail,
		},
	}
}

func (s *resendMailer) Send(templateFile string, receipient EmailData, data any, isSandbox bool) (string, error) {
	if isSandbox {
		return "SANDBOX_ID", nil
	}

	// template parsing and building
	tmpl, err := template.ParseFS(FS, fmt.Sprintf("templates/%s", templateFile))
	if err != nil {
		return "", err
	}

	subject := new(bytes.Buffer)
	if err = tmpl.ExecuteTemplate(subject, "subject", data); err != nil {
		return "", err
	}

	body := new(bytes.Buffer)
	if err = tmpl.ExecuteTemplate(body, "body", data); err != nil {
		return "", err
	}

	message := &resend.SendEmailRequest{
		From:    s.sender.Format(),
		To:      []string{receipient.Format()},
		Subject: subject.String(),
		Html:    body.String(),
	}

	var sendErrs []error
	for i := 1; i <= maxRetries; i++ {
		resp, err := s.client.Emails.Send(message)
		if err != nil {
			sendErrs = append(sendErrs, err)

			time.Sleep(time.Second * time.Duration(i*2))
			continue
		}
		return resp.Id, nil
	}

	return "", NewSendError(sendErrs, "failed to send email to '%s' after %d attempts", receipient.Format(), maxRetries)
}