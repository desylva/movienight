package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/desylva/movienight/models"
	"github.com/gobuffalo/pop"
	"github.com/pkg/errors"
)

// Allocate an empty User
var user = &models.User{}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	// Get the Session from the context
	s := c.Session()

	// Find User ID from the parameter user_id.
	var userID string
	userID = c.Param("user_id")
	if len(userID) > 1 {
		setUserFromID(userID, c)
	} else {
	// Find the User ID from the session user_id
		su := s.Get("user_id")
		userID, ok := su.(string)
		if ok {
			setUserFromID(userID, c)
		}
	}

	c.Set("user", user)
	return c.Render(200, r.HTML("home.html"))
}

func setUserFromID(uuid string, c buffalo.Context) error {
	// Get the Session from the context
	s := c.Session()

	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return errors.WithStack(errors.New("no transaction found"))
	}
	
	// Retrieve user from db using user ID
	if err := tx.Find(user, uuid); err != nil {
		return c.Error(404, err)
	}

	s.Clear()
	s.Set("user_id", user.ID.String())
	c.Set("user", user)

	return nil
}