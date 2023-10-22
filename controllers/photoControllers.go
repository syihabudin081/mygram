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



//For Upload Images

// func CreatePhoto(c *gin.Context) {
// 	db := database.GetDB()
// 	userData := c.MustGet("userData").(jwt.MapClaims)
// 	contentType := helpers.GetContentType(c)

// 	Photo := models.Photo{}
// 	userID := uint(userData["id"].(float64))

// 	if contentType != appJSON {
// 		c.ShouldBindJSON(&Photo)
// 	} else {
// 		c.ShouldBind(&Photo)
// 	}

// 	// Dapatkan file gambar dari request form
// 	file, err := c.FormFile("image")
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Bad Request",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// Simpan gambar ke penyimpanan awan (contoh: Amazon S3)
// 	imageURL, err := saveToCloudStorage(file)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{
// 			"error": "Internal Server Error",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	// Simpan URL gambar ke dalam basis data
// 	Photo.UserID = userID
// 	Photo.ImageURL = imageURL

// 	err = db.Debug().Create(&Photo).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"error": "Bad Request",
// 			"message": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, Photo)
// }

// // Fungsi untuk menyimpan gambar ke penyimpanan awan (contoh: Amazon S3)
// func saveToCloudStorage(file *multipart.FileHeader) (string, error) {
// 	// Implementasikan logika untuk menyimpan gambar ke penyimpanan awan disini
// 	// Anda perlu mengganti ini sesuai dengan penyimpanan awan yang Anda gunakan
// 	// Dapatkan URL gambar setelah berhasil diunggah
// 	// Atau kembalikan kesalahan jika terjadi masalah
// 	// Contoh: Simpan ke Amazon S3 dan dapatkan URL gambar
// 	// Replace this with your own implementation
// 	// Example: Save to Amazon S3 and get the image URL
// 	imageURL := "https://your-s3-bucket.s3.amazonaws.com/" + file.Filename
// 	return imageURL, nil
// }