package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/pop/slices"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Movie struct {
	ID           uuid.UUID    `json:"id" db:"id"`
	CreatedAt    time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time    `json:"updated_at" db:"updated_at"`
	Name         string       `json:"name" db:"name"`
	UserUUID     uuid.UUID    `json:"user_uuid" db:"user_uuid"`
	Imdb         nulls.String `json:"imdb" db:"imdb"`
	UsersFor     slices.UUID  `json:"users_for" db:"users_for"`
	UsersAgainst slices.UUID  `json:"users_against" db:"users_against"`
}

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
	ImdbID     string `json:"imdbId"`
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
