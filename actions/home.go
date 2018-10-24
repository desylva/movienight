package actions

import (
	// "github.com/desylva/movienight/models"
	"github.com/gobuffalo/buffalo"
	// "github.com/gobuffalo/pop"
	// "github.com/pkg/errors"
	// "os/user"
)

// Allocate an empty User
// var user = &models.User{}

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {

	// 	// Get the Session from the context
	// 	s := c.Session()

	// 	// Find User ID from the parameter user_id.
	// 	var userID string
	// 	userID = c.Param("user_id")
	// 	if len(userID) > 1 {
	// 		setUserFromID(userID, c)
	// 	} else {
	// 		// Find the User ID from the session user_id
	// 		var user models.User
	// 		user, err = s.Get("current_user_id")
	// 		if err != nil {

	// 		}
	// 	}

	return c.Render(200, r.HTML("home.html"))
}

// func setUserFromID(uuid string, c buffalo.Context) error {
// 	// Get the Session from the context
// 	s := c.Session()

// 	// Get the DB connection from the context
// 	tx, ok := c.Value("tx").(*pop.Connection)
// 	if !ok {
// 		return errors.WithStack(errors.New("no transaction found"))
// 	}

// 	// Retrieve user from db using user ID
// 	if err := tx.Find(user, uuid); err != nil {
// 		s.Clear()
// 		return c.Error(404, err)
// 	}

// 	s.Clear()
// 	s.Set("current_user", user.ID.String())
// 	c.Set("user", user)

// 	return nil
// }
