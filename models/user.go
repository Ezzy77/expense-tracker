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
