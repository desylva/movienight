package mailers

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/buffalo/mail"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/envy"
	"github.com/pkg/errors"
)

// SendAccountReset submit an welcome email to user
func SendAccountReset(c buffalo.Context, emails []string, hash string) error {
	m := mail.New(c)

	siteURL := envy.Get("SITE_URL", "")
	emailFrom := envy.Get("EMAIL_FROM", "")

	URIRecoverAccount := siteURL + "/users/recover/" + hash + "/"
	data := render.Data{
		"URIRecoverAccount": URIRecoverAccount,
	}

	// fill in with your stuff:
	m.Subject = "Account Password Reset"
	m.From = emailFrom
	m.To = emails
	err := m.AddBody(r.HTML("account_reset.html"), data)
	if err != nil {
		return errors.WithStack(err)
	}
	return smtp.Send(m)
}
