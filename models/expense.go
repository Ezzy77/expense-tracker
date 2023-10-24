package models

import (
	"database/sql"
	"errors"
	"fmt"
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

type ExpenseModel struct {
	DB *sql.DB
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

func (e *ExpenseModel) CreateExpense(exp *Expense) error {

	stmt := ` INSERT INTO expense
	(id, item, price, store, date)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := e.DB.Query(
		stmt,
		exp.ID,
		exp.Item,
		exp.Price,
		exp.Store,
		exp.Date,
	)

	if err != nil {
		fmt.Println("me now ", err)
		return err
	}

	return nil
}

func (e *ExpenseModel) GetExpenses() ([]*Expense, error) {
	rows, err := e.DB.Query("SELECT * FROM expense")
	if err != nil {
		return nil, err
	}

	var expenses []*Expense

	for rows.Next() {
		expense := Expense{}
		if err := rows.Scan(
			&expense.ID,
			&expense.Item,
			&expense.Price,
			&expense.Store,
			&expense.Date,
		); err != nil {
			return nil, err
		}
		expenses = append(expenses, &expense)
	}

	return expenses, nil

}

func (e *ExpenseModel) GetExpenseById(id string) (*Expense, error) {
	stmt := `SELECT * FROM expense WHERE id = $1`

	expense := Expense{}

	err := e.DB.QueryRow(stmt, id).Scan(
		&expense.ID,
		&expense.Item,
		&expense.Price,
		&expense.Store,
		&expense.Date,
	)
	if err != nil {
		return nil, err
	}

	return &expense, nil
}

func (e *ExpenseModel) DeleteExpense(id string) error {
	stmt := `DELETE FROM expense WHERE id = $1`

	result, err := e.DB.Exec(stmt, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("record not found")
	}

	return nil
}

func (e *ExpenseModel) UpdateExpense(id string, exp *Expense) (*Expense, error) {
	stmt := `UPDATE expense SET item=$1, price=$2, store=$3
	 WHERE id = $4`

	expense := Expense{
		ID:    exp.ID,
		Item:  exp.Item,
		Price: exp.Price,
		Store: exp.Store,
		Date:  exp.Date,
	}

	updatedExp := Expense{}

	row, err := e.DB.Query(stmt, expense.Item, expense.Price, expense.Store, id)
	if err != nil {
		return nil, err
	}

	row.Scan(
		&updatedExp.ID,
		&updatedExp.Item,
		&updatedExp.Price,
		&updatedExp.Store,
		&updatedExp.Date,
	)
	fmt.Println("=====> ", updatedExp)

	return &updatedExp, nil
}
