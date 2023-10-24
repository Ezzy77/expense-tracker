package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) userSignUpHandler(c *gin.Context) {
	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func (app *application) userLoginHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user loged in",
	})
}

func (app *application) userLogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user logged out",
	})
}
