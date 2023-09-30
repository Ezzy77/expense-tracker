package main

import (
	"net/http"

	"github.com/ezzy77/expense-tracker/internals/models"
	"github.com/gin-gonic/gin"
)

var expenses []models.Expense

func (app *application) getExpensesHandler(c *gin.Context) {

	c.JSON(http.StatusOK, expenses)

}

func (app *application) createExpensesHandler(c *gin.Context) {

	var expense models.Expense

	err := c.ShouldBindJSON(&expense)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "could not parse json",
		})
	}

	expenses = append(expenses, expense)
	c.JSON(http.StatusAccepted, expense)

}

func (app *application) getExpenseHandler(c *gin.Context) {

}

func (app *application) updateExpenseHandler(c *gin.Context) {

}

func (app *application) deleteExpenseHandler(c *gin.Context) {

}
