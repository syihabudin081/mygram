package controllers

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"mygram/database"
	"mygram/helpers"
	"mygram/models"
	"net/http"
	"os"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"

	"cloud.google.com/go/storage"
	"github.com/google/uuid"
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

// func CreatePhoto(c *gin.Context) {
// 	db := database.GetDB()
// 	userData := c.MustGet("userData").(jwt.MapClaims)
// 	contentType := helpers.GetContentType(c)
// 	Photo := models.Photo{}
// 	useriD := userData["id"].(float64)

// 	if contentType != appJSON {
// 		c.ShouldBindJSON(&Photo)
// 	} else {
// 		c.ShouldBind(&Photo)
// 	}

// 	Photo.UserID = uint(useriD)
// 	err := db.Debug().Create(&Photo).Error

// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"err": "Bad Request",

// 			"message": err.Error(),
// 		})
// 		return

// 	}

// 	c.JSON(http.StatusOK, Photo)

// }

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

	err := db.Debug().Preload("User").Find(&Photos).Error

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

	photoID, _ := strconv.Atoi(c.Param("photoId"))

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

	photoID, _ := strconv.Atoi(c.Param("photoId"))

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

func CreatePhoto(c *gin.Context) {
    db := database.GetDB()
    userData := c.MustGet("userData").(jwt.MapClaims)
    contentType := helpers.GetContentType(c)

    // Dapatkan file gambar dari request form
	
    f, err := c.FormFile("file_input")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error": "file_input",
        })
		
        return
    }

    fmt.Print("content type", contentType)
    Photo := models.Photo{}
    userID := uint(userData["id"].(float64))

    // Ambil data dari form-data
   
    caption := c.PostForm("caption")

    // Isi struktur Photo dengan data dari form-data

    Photo.Caption = caption

    imageURL, err := saveToCloudStorage(f)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "error":   "Internal Server Error",
            "message": err.Error(),
        })
        return
    }

    // Simpan URL gambar ke dalam basis data
    Photo.UserID = userID
    Photo.Photo_URL = imageURL

    err = db.Debug().Create(&Photo).Error

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, Photo)
}


// // Fungsi untuk menyimpan gambar ke penyimpanan awan (contoh: Amazon S3)

func saveToCloudStorage(file *multipart.FileHeader) (string, error) {
	// Your Google Cloud Storage bucket name and service account JSON key file
	bucketName := os.Getenv("BUCKET_NAME")
	keyFile := os.Getenv("SERVICE_ACCOUNT_FILE")

	// Create a context and a client using your service account key
	ctx := context.Background()
	client, err := storage.NewClient(ctx, option.WithServiceAccountFile(keyFile))
	if err != nil {
		return "", err
	}
	defer client.Close()

	// Generate a unique file name for the uploaded file
	fileName := uuid.New().String() + ".jpg" // You can change the file extension as needed

	// Open the file to be uploaded
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Create an object handle for your Google Cloud Storage bucket
	bucket := client.Bucket(bucketName)
	obj := bucket.Object(fileName)

	// Create an object writer
	wc := obj.NewWriter(ctx)

	// Copy the contents of the file to the object writer
	if _, err := io.Copy(wc, src); err != nil {
		return "", err
	}

	// Close the object writer to flush the data to Google Cloud Storage
	if err := wc.Close(); err != nil {
		return "", err
	}

	// Public URL for the uploaded file
	url := fmt.Sprintf("https://storage.googleapis.com/%s/%s", bucketName, fileName)

	return url, nil
}
