package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

// SendAccountWelcomes submit an welcome email to user
func SendAccountWelcomes(emails []string) error {
	m := mail.NewMessage()

	emailFrom := envy.Get("EMAIL_FROM", "")

	// fill in with your stuff:
	m.Subject = "Welcome"
	m.From = emailFrom
	m.To = emails
	err := m.AddBody(r.HTML("account_welcome.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
