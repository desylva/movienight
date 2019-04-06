package models

import (
	"encoding/json"
	"errors"
	"math/rand"
	"strings"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	// "github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/lucasb-eyer/go-colorful"
	"golang.org/x/crypto/bcrypt"
)

// User structure matching db
type User struct {
	ID                uuid.UUID `json:"id" db:"id"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
	Name              string    `json:"name" db:"name"`
	Email             string    `json:"email" db:"email"`
	Active            bool      `json:"active" db:"active"`
	Color             string    `json:"color" db:"color"`
	PasswordHash      string    `json:"-" db:"password_hash"`
	PasswordHashReset string    `json:"-" db:"password_hash_reset"`
	Verified          bool      `json:"verified" db:"verified"`
	VerificationHash  string    `json:"-" db:"verification_hash"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
		&validators.StringIsPresent{Field: u.Email, Name: "Email"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// UserColorGenerator generates a random Hex color.
func UserColorGenerator() string {
	rand.Seed(time.Now().UTC().UnixNano())
	c := colorful.HappyColor()
	return c.Hex()
}

// GetByEmail retrieve a user from the DB using its email address
func (u *User) GetByEmail(c buffalo.Context, email string) (User, error) {
	// Get the DB connection from the context
	tx, ok := c.Value("tx").(*pop.Connection)
	if !ok {
		return User{}, errors.New("no connection found")
	}

	// find a user with the email
	if err := tx.Where("email = ?", strings.ToLower(email)).First(u); err != nil {
		return User{}, errors.New("no user found")
	}

	return *u, nil
}

// EncryptPassword makes a hash from the string password
func EncryptPassword(p string) string {
	pwd, err := bcrypt.GenerateFromPassword([]byte(strings.TrimSpace(p)), 8)

	if err != nil {
		panic("could not encrypt password")
	}

	return string(pwd)
}

// ComparePassword compares the hash of password to a password string
func ComparePassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err != nil {
		return false
	}
	return true
}
