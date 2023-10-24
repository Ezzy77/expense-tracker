package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes() *gin.Engine {

	route := gin.Default()

	//expense routes
	route.GET("/v1/api/healthcheck", app.healthcheckHandler)
	route.GET("/v1/api/expenses", app.getExpensesHandler)
	route.POST("/v1/api/expenses", app.createExpensesHandler)
	route.GET("/v1/api/expenses/:id", app.getExpenseHandler)
	route.DELETE("/v1/api/expenses/:id", app.deleteExpenseHandler)
	route.PATCH("/v1/api/expenses/:id", app.updateExpenseHandler)

	//user routes
	route.POST("/v1/api/user/signup", app.userSignUpHandler)
	route.POST("/v1/api/user/login", app.userLoginHandler)
	route.POST("/v1/api/user/logout", app.userLogoutHandler)

	return route
}
