package models

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	FirstName string
	LastName  string
	Email     string
	Password  string
	Date      time.Time
}

type UserModel struct {
	DB *sql.DB
}

func NewUser(firstName, lastName, email, password string) *User {
	return &User{
		ID:        uuid.New(),
		FirstName: firstName,
		LastName:  lastName,
		Email:     email,
		Password:  password,
		Date:      time.Now().UTC(),
	}
}

func (u *UserModel) Insert(firstName, lastName, email, password string) error {

	return nil
}

func (u *UserModel) Authenticate(email, password string) (int, error) {
	return 0, nil
}

func (u *UserModel) Exists(id int) (bool, error) {
	return false, nil
}
