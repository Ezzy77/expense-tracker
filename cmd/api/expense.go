package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ezzy77/expense-tracker/models"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var expenses []models.Expense

func (app *application) getExpensesHandler(c *gin.Context) {

	res, err := app.store.GetExpenses()
	if err != nil {
		fmt.Println(err)
	}

	c.JSON(202, res)

}

func (app *application) createExpensesHandler(c *gin.Context) {

	var expense models.Expense

	err := c.ShouldBindJSON(&expense)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse json",
		})
	}

	expense.ID = uuid.New()
	expense.Date = time.Now()

	validate := validator.New()
	err = validate.Struct(expense)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "error validating",
		})
		fmt.Println(err)
		return
	}

	err = app.store.CreateExpense(&expense)
	if err != nil {
		fmt.Print(err)
	}

	//expenses = append(expenses, expense)
	c.JSON(http.StatusAccepted, expense)

}

func (app *application) getExpenseHandler(c *gin.Context) {

}

func (app *application) updateExpenseHandler(c *gin.Context) {

}

func (app *application) deleteExpenseHandler(c *gin.Context) {

}
