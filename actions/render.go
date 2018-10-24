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

//{"Title":"Johnny English Strikes Again","Year":"2018","Rated":"PG",
//"Released":"26 Oct 2018","Runtime":"88 min",
//"Genre":"Action, Adventure, Comedy",
//"Director":"David Kerr","Writer":"William Davies (screenplay by)",
//"Actors":"Olga Kurylenko, Rowan Atkinson, Emma Thompson, Charles Dance",
//"Plot":"After a cyber-attack reveals the identity of all of the active undercover agents
//in Britain, Johnny English is forced to come out of retirement to find the mastermind hacker.",
//"Language":"English","Country":"UK, France, USA","Awards":"N/A",
//"Poster":"https://m.media-amazon.com/images/M/MV5BMjI4M
//jQ3MjI5MV5BMl5BanBnXkFtZTgwNjczMDE4NTM@._V1_SX300.jpg",
//"Ratings":[{"Source":"Internet Movie Database","Value":"6.6/10"},
//{"Source":"Rotten Tomatoes","Value":"37%"},{"Source":"Metacritic","Value":"35/100"}],
//"Metascore":"35","imdbRating":"6.6","imdbVotes":"9,960","imdbID":"tt6921996",
//"Type":"movie","DVD":"N/A","BoxOffice":"N/A","Production":"Universal Pictures",
//"Website":"http://www.johnnyenglishmovie.com/","Response":"True"}
