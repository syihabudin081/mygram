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

// CreatePhoto creates a new photo
// @Summary Create a new photo
// @Description Create a new photo
// @Tags photos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param photo body models.Photo true "Photo"
// @Success 200 {object} models.Photo
// @Failure 400 
// @Router /photos [post]
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

// GetPhotos gets all photos
// @Summary Get all photos
// @Description Get all photos
// @Tags photos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Success 200 {object} models.Photo
// @Failure 400 
// @Router /photos [get]
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

// GetPhotoById gets a photo by id
// @Summary Get a photo by id
// @Description Get a photo by id
// @Tags photos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param id path int true "Photo ID"
// @Success 200 {object} models.Photo
// @Failure 400 
// @Router /photos/{id} [get]
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

// UpdatePhoto updates a photo
// @Summary Update a photo
// @Description Update a photo
// @Tags photos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param id path int true "Photo ID"
// @Param photo body models.Photo true "Photo"
// @Success 200 {object} models.Photo
// @Failure 400 
// @Router /photos/{id} [put]
func UpdatePhoto(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)
	Photo := models.Photo{}

	photoID,_ := strconv.Atoi(c.Param("photoId"))

	userID := uint(userData["id"].(float64)) 
	if contentType != appJSON {
		c.ShouldBindJSON(&Photo)
	} else {
		c.ShouldBind(&Photo)
	}

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

// DeletePhotoByID deletes a photo by id
// @Summary Delete a photo by id
// @Description Delete a photo by id
// @Tags photos
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param id path int true "Photo ID"
// @Success 200 {object} models.Photo
// @Failure 400 
// @Router /photos/{id} [delete]
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