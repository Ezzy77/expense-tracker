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

func (app *application) getExpensesHandler(c *gin.Context) {

	res, err := app.expenses.GetExpenses()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	c.JSON(http.StatusAccepted, res)

}

func (app *application) createExpensesHandler(c *gin.Context) {

	var expense models.Expense

	err := c.ShouldBindJSON(&expense)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse json",
		})
		return
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

	err = app.expenses.CreateExpense(&expense)
	if err != nil {
		fmt.Print(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "bad request",
		})
		return
	}

	c.JSON(http.StatusCreated, &expense)

}

func (app *application) getExpenseHandler(c *gin.Context) {
	id := c.Param("id")

	expense, err := app.expenses.GetExpenseById(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "could not find expense",
		})
		return
	}

	c.JSON(http.StatusAccepted, expense)

}

func (app *application) updateExpenseHandler(c *gin.Context) {
	id := c.Param("id")

	expense := models.Expense{}

	err := c.ShouldBindJSON(&expense)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse json",
		})
		return
	}

	res, err := app.expenses.UpdateExpense(id, &expense)
	if err != nil {
		fmt.Println(err)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not update json",
		})
		return
	}

	c.JSON(http.StatusAccepted, res)

}

func (app *application) deleteExpenseHandler(c *gin.Context) {
	id := c.Param("id")

	err := app.expenses.DeleteExpense(id)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "could not find expense",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "deleted successfuly",
	})

}
