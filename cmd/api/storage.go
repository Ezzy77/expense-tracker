package main

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/ezzy77/expense-tracker/models"
	_ "github.com/lib/pq"
)

type Storage interface {
	CreateExpense(exp *models.Expense) error
	GetExpenses() ([]*models.Expense, error)
	GetExpenseById(id string) (*models.Expense, error)
	DeleteExpense(id string) error
	UpdateExpense(id string, exp *models.Expense) (*models.Expense, error)
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
		fmt.Println("me now ", err)
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

func (s *PostgresStore) GetExpenseById(id string) (*models.Expense, error) {
	stmt := `SELECT * FROM expense WHERE id = $1`

	expense := models.Expense{}

	err := s.db.QueryRow(stmt, id).Scan(
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

func (s *PostgresStore) DeleteExpense(id string) error {
	stmt := `DELETE FROM expense WHERE id = $1`

	result, err := s.db.Exec(stmt, id)
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

func (s *PostgresStore) UpdateExpense(id string, exp *models.Expense) (*models.Expense, error) {
	stmt := `UPDATE expense SET item=$1, price=$2, store=$3
	 WHERE id = $4`

	expense := models.Expense{
		ID:    exp.ID,
		Item:  exp.Item,
		Price: exp.Price,
		Store: exp.Store,
		Date:  exp.Date,
	}

	updatedExp := models.Expense{}

	row, err := s.db.Query(stmt, expense.Item, expense.Price, expense.Store, id)
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
