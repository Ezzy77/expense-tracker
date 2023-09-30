package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) healthcheckHandler(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"Hi": "there",
	})

}
