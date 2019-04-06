package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

func SendAccountWelcomes(emails []string) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Account Welcome"
	m.From = "movieparty32@gmail.com"
	m.To = emails
	err := m.AddBody(r.HTML("account_welcome.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
