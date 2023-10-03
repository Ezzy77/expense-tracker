package models

import (
	"time"

	_ "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Expense struct {
	ID    uuid.UUID `json:"id" `
	Item  string    `json:"item" validate:"required,max=20,min=2"`
	Price float64   `json:"price" validate:"required,number"`
	Store string    `json:"store"`
	Date  time.Time `json:"date"`
}

func NewExpense(item string, price float64, store string) *Expense {
	return &Expense{
		ID:    uuid.New(),
		Item:  item,
		Price: price,
		Store: store,
		Date:  time.Now().UTC(),
	}
}
