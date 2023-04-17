package controllers

import (
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func CreatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Photo := models.Photo{}
	useriD := userData["id"].(float64)

	if contentType != appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

	Photo.UserID = uint(useriD)
	err := db.Debug().Create(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, Photo)

}

func GetPhotos(c *gin.Context) {
	db := database.GetDB()
	Photos := []models.Photo{}

	err := db.Debug().Find(&Photos).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, Photos)

}

func GetPhotoById(c *gin.Context) {
	db := database.GetDB()
	Photo := models.Photo{}
	id := c.Param("id")

	err := db.Debug().Where("id = ?", id).First(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, Photo)

}

func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoID,_ := strconv.Atoi(c.Param("photoId"))

	userID := uint(userData["id"].(float64)) 

	Photo.UserID = userID
	Photo.ID = uint(photoID)

	err := db.Debug().Where("id = ?", photoID).Updates(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo Updated",
	})
}

func DeletePhotoByID(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	Photo := models.Photo{}

	photoID,_ := strconv.Atoi(c.Param("photoId"))

	userID := uint(userData["id"].(float64)) 

	Photo.UserID = userID
	Photo.ID = uint(photoID)

	err := db.Where("id = ?", photoID).Delete(&Photo).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Photo Deleted",
	})


}