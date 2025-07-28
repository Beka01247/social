package mailer

import "embed"

const (
	FromName            = "SocialApp"
	maxRetires          = 3
	UserWelcomeTemplate = "user_invitation.tmpl"
)

//go:embed "template"
var FS embed.FS

type Client interface {
	Send(templateFile, username, email string, data any, isSandbox bool) (int, error)
}
