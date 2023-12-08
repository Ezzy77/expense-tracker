package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ezzy77/expense-tracker/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (app *application) userSignUpHandler(c *gin.Context) {

	var user models.User

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	user.ID = uuid.New()
	user.Date = time.Now()

	err = app.users.Insert(user)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created",
	})
}

func (app *application) userLoginHandler(c *gin.Context) {

	user := models.User{}

	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error with data",
		})
	}

	sessKey, err := app.users.Authenticate(user.Email, user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session key": sessKey,
	})
}

func (app *application) userLogoutHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "user logged out",
	})
}

func (app *application) userListHandler(c *gin.Context) {
	res, err := app.users.GetAll()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"error": "not found",
		})
		return
	}

	c.JSON(http.StatusAccepted, res)
}
