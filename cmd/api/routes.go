package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {

	route := gin.Default()

	route.GET("/v1/api/healthcheck", app.healthcheckHandler)
	route.GET("/v1/api/expenses", app.getExpensesHandler)
	route.POST("/v1/api/expenses", app.createExpensesHandler)
	route.GET("/v1/api/expenses/:id", app.getExpenseHandler)
	route.DELETE("/v1/api/expenses", app.deleteExpenseHandler)
	route.PATCH("/v1/api/expenses", app.updateExpenseHandler)

	return route
}
