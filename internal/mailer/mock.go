package mailer

import (
	"fmt"
	"log"
)

type mockMailer struct {
	fromEmail string
}

func NewMockMailer(fromEmail string) *mockMailer {
	return &mockMailer{
		fromEmail: fromEmail,
	}
}

func (m *mockMailer) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	log.Printf("MOCK MAILER: Email would be sent to %s <%s>", username, email)
	log.Printf("MOCK MAILER: Template: %s", templateFile)
	log.Printf("MOCK MAILER: Sandbox: %v", isSandbox)
	log.Printf("MOCK MAILER: Data: %+v", data)

	fmt.Printf("✅ Mock email sent successfully to %s\n", email)
	return 200, nil
}
