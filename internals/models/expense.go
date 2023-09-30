package models

import "time"

type Expense struct {
	ID        int
	Item      string
	Date      time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}
