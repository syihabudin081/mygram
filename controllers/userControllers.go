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

// UserRegister registers a new user
// @Summary Register a new user
// @Description Create a new user with the given details
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 
// @Router /users/register [post]
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

// UserLogin logs in a user
// @Summary Login a user
// @Description Log in a user with the given credentials
// @Tags Users
// @Accept json
// @Produce json
// @Success 200 
// @Router /users/login [post]
func UserLogin(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)
	_,_ = db, contentType
	User := models.User{}
	password := ""

	if contentType != appJSON {
		c.ShouldBindJSON(&User)
	} else {
		c.ShouldBind(&User)
	}

	password = User.Password
	err := db.Debug().Where("email = ?", User.Email).Take(&User).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "User not found",
			"error":   err.Error(),
		})
		return
	}

	comparePass := helpers.ComparePass([]byte(User.Password), []byte(password))

	if !comparePass {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status":  http.StatusUnauthorized,
			"message": "Wrong password",
		})
		return
	}

	token := helpers.GenerateToken(User.ID, User.Email, User.Username)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Login success",
		"token":   token,
	})



}