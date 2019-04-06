package mailers

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/buffalo/mail"
	"github.com/pkg/errors"
)

func SendAccountResets() error {
	m := mail.NewMessage()

	// fill in with your stuff:
	m.Subject = "Account Reset"
	m.From = ""
	m.To = []string{}
	err := m.AddBody(r.HTML("account_reset.html"), render.Data{})
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
