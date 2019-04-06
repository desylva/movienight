package mailers

import (
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/pkg/errors"
)

// SendAccountConfirms sends an email to confirm account creation
func SendAccountConfirms(emails []string) error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Account Confirm"
	m.From = ""
	m.To = emails
	err := m.AddBody(r.HTML("account_confirm.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
