package main

import (
	"database/sql"
	"fmt"

	"github.com/ezzy77/expense-tracker/models"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateExpense(*models.Expense) error
	GetExpenses() ([]*models.Expense, error)
	DeleteExpense(int) error
	UpdateExpense(int) error
	GetExpenseById(int) error
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	connStr := "user=postgres dbname=expense_tracker password=mysecretpassword sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if err := db.Ping(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) init() error {

	query := `create table if not exists expense(
		id UUID primary key,
		item TEXT NOT NULL,
		price INTEGER NOT NULL,
		store TEXT NOT NULL,
		date TIMESTAMP
	)`

	_, err := s.db.Exec(query)
	return err

}

func (s *PostgresStore) CreateExpense(exp *models.Expense) error {

	stmt := ` INSERT INTO expense
	(id, item, price, store, date)
	VALUES ($1, $2, $3, $4, $5)`

	_, err := s.db.Query(
		stmt,
		exp.ID,
		exp.Item,
		exp.Price,
		exp.Store,
		exp.Date,
	)
	if err != nil {
		return err
	}

	return nil
}

func (s *PostgresStore) GetExpenses() ([]*models.Expense, error) {
	rows, err := s.db.Query("SELECT * FROM expense")
	if err != nil {
		return nil, err
	}

	var expenses []*models.Expense

	for rows.Next() {
		expense := models.Expense{}
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

func (s *PostgresStore) DeleteExpense(id int) error {
	return nil
}
func (s *PostgresStore) UpdateExpense(id int) error {
	return nil
}

func (s *PostgresStore) GetExpenseById(id int) error {
	return nil
}
