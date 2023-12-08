package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID        uuid.UUID
	FirstName string    `json:"first_name" `
	LastName  string    `json:"last_name" `
	Email     string    `json:"email" `
	Password  string    `json:"password" `
	Date      time.Time `json:"date" `
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

func (u *UserModel) Insert(user User) error {

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return err
	}

	stmt := ` INSERT INTO client 
	(id, first_name, last_name, email, hashed_password, date)
	VALUES ($1, $2, $3, $4, $5, $6)`

	_, err = u.DB.Exec(stmt,
		user.ID,
		user.FirstName,
		user.LastName,
		user.Email,
		string(hashedPassword),
		user.Date,
	)

	if err != nil {
		return err
	}

	return nil
}

func (u *UserModel) Authenticate(email, password string) (uuid.UUID, error) {
	var id uuid.UUID
	var hashedPassword []byte

	stmt := `SELECT id, hashed_password  FROM client WHERE email = $1`

	err := u.DB.QueryRow(stmt, email).Scan(
		&id,
		&hashedPassword,
	)
	if err != nil {
		return uuid.Nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
	if err != nil {
		return uuid.Nil, err
	}

	return id, nil
}

func (u *UserModel) Exists(id int) (bool, error) {
	return false, nil
}

func (u *UserModel) GetAll() ([]*User, error) {

	rows, err := u.DB.Query("SELECT * FROM client")
	if err != nil {
		return nil, err
	}

	var users []*User

	for rows.Next() {
		user := User{}
		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.Email,
			&user.Password,
			&user.Date,
		); err != nil {
			return nil, err
		}
		users = append(users, &user)
	}

	return users, nil
}
