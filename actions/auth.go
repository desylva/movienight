package actions

import (
	// "github.com/desylva/movienight/helpers"
	// "github.com/desylva/movienight/models"
	"github.com/gobuffalo/buffalo"
)

// AuthResource is the resource for the Movie model
type AuthResource struct {
	buffalo.Resource
}

// AuthLogin default implementation.
func AuthLogin(c buffalo.Context) error {
	return c.Render(200, r.HTML("auth/login.html"))
}

// AuthNew default implementation.
func AuthNew(c buffalo.Context) error {
	c.Set("email", c.Param("email"))
	c.Set("password", c.Param("password"))
	return c.Render(200, r.HTML("auth/new.html"))
}
