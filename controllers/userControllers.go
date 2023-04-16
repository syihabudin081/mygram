package controllers

import (
	"mygram/database"
	"mygram/models"
	"net/http"
	"mygram/helpers"
	"github.com/gin-gonic/gin"
)

var (
	appJSON = "application/json"
)


func UserRegister(c *gin.Context) {

	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_, _ = db, contentType
	User := models.User{}

	if contentType != appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	err := db.Debug().Create(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User not created",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":    http.StatusOK,
		"id":        User.ID,
		"email":     User.Email,
		"username": User.Username,
		"age":       User.Age,
	})

}