package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Movie struct {
	ID           uuid.UUID   `json:"id" db:"id"`
	CreatedAt    time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time   `json:"updated_at" db:"updated_at"`
	Name         string      `json:"name" db:"name"`
	UserUUID     uuid.UUID   `json:"user_uuid" db:"user_uuid"`
	ImdbID       string      `json:"imdb_id" db:"imdb_id"`
	UsersFor     slices.UUID `json:"users_for" db:"users_for"`
	UsersAgainst slices.UUID `json:"users_against" db:"users_against"`
	Score        int         `json:"score" db:"score"`
}

// Example of the response from a request to the OMDB API:
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
type ImdbData struct {
	Title      string `json:"title"`
	Year       string `json:"year"`
	Rated      string `json:"rated"`
	Released   string `json:"released"`
	Runtime    string `json:"runtime"`
	Genre      string `json:"genre"`
	Director   string `json:"director"`
	Writer     string `json:"writer"`
	Actors     string `json:"actors"`
	Plot       string `json:"plot"`
	Poster     string `json:"poster"`
	ImdbID     string `json:"imdbID"`
	Production string `json:"production"`
	Website    string `json:"website"`
}

// String is not required by pop and may be deleted
func (m Movie) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Movies is not required by pop and may be deleted
type Movies []Movie

// String is not required by pop and may be deleted
func (m Movies) String() string {
	jm, _ := json.Marshal(m)
	return string(jm)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (m *Movie) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: m.Name, Name: "Name"},
		&validators.StringIsPresent{Field: m.ImdbID, Name: "ImdbID"},
		//&validators.StringIsPresent{Field: m.User, Name: "User"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (m *Movie) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (m *Movie) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// UserStringToUserUUID parses a string uuid into a UUID uuid.
func (m *Movie) UserStringToUserUUID(uu string) (uuid.UUID, error) {
	// Parse a UUID from a string.
	u, err := uuid.FromString(uu)
	if err != nil {
		return u, err
	}
	return u, nil
}

func (m Movie) UserIsFor(id uuid.UUID) bool {
	for _, usr := range m.UsersFor {
		if usr == id {
			return true
		}
	}
	return false
}

func (m *Movie) AddUserFor(id uuid.UUID) {
	set := make(map[uuid.UUID]struct{})
	set[id] = struct{}{}
	for _, user := range m.UsersFor {
		set[user] = struct{}{}
	}
	users := make(slices.UUID, 0, len(set))
	for user := range set {
		users = append(users, user)
	}
	m.UsersFor = users
}

func (m *Movie) RemoveUserFor(id uuid.UUID) {
	set := make(map[uuid.UUID]struct{})
	for _, user := range m.UsersFor {
		if user != id {
			set[user] = struct{}{}
		}
	}
	users := make(slices.UUID, 0, len(set))
	for user := range set {
		users = append(users, user)
	}
	m.UsersFor = users
}

func (m Movie) UserIsAgainst(id uuid.UUID) bool {
	for _, usr := range m.UsersAgainst {
		if usr == id {
			return true
		}
	}
	return false
}

func (m *Movie) AddUserAgainst(id uuid.UUID) {
	set := make(map[uuid.UUID]struct{})
	set[id] = struct{}{}
	for _, user := range m.UsersAgainst {
		set[user] = struct{}{}
	}
	users := make(slices.UUID, 0, len(set))
	for user := range set {
		users = append(users, user)
	}
	m.UsersAgainst = users
}

func (m *Movie) RemoveUserAgainst(id uuid.UUID) {
	set := make(map[uuid.UUID]struct{})
	for _, user := range m.UsersAgainst {
		if user != id {
			set[user] = struct{}{}
		}
	}
	users := make(slices.UUID, 0, len(set))
	for user := range set {
		users = append(users, user)
	}
	m.UsersAgainst = users
}
