package actions

import (
	"github.com/desylva/movienight/models"
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
	"github.com/gobuffalo/plush"
	"github.com/gobuffalo/pop"
)

var r *render.Engine
var assetsBox = packr.NewBox("../public")

func init() {
	r = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../templates"),
		AssetsBox:    assetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{
			// uncomment for non-Bootstrap form helpers:
			// "form":     plush.FormHelper,
			// "form_for": plush.FormForHelper,
			"getUserName": func(uuid string, help plush.HelperContext) string {
				// Get the DB connection from the context
				tx, ok := help.Value("tx").(*pop.Connection)
				if !ok {
					return ""
				}
				// Allocate an empty User
				user := &models.User{}

				// To find the User the parameter user_id is used.
				if err := tx.Find(user, uuid); err != nil {
					return ""
				}

				return user.Name
			},
		},
	})
}
