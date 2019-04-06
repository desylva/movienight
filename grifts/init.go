package grifts

import (
	"github.com/desylva/movieparty/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
