package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

type Movie struct {
	ID        uuid.UUID    `json:"id" db:"id"`
	CreatedAt time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt time.Time    `json:"updated_at" db:"updated_at"`
	Name      string       `json:"name" db:"name"`
	UserUUID  string       `json:"user_uuid" db:"user_uuid"`
	Imdb      nulls.String `json:"imdb" db:"imdb"`
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
