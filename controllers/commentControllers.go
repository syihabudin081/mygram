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

func CreateComment(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	useriD := userData["id"].(float64)
	photoID, err := strconv.Atoi(c.Param("photoId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid photo ID",
		})
		return
	}

	// Check if photo with given ID exists
	var photo models.Photo
	err = db.Where("id = ?", photoID).First(&photo).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Photo with given ID does not exist",
		})
		return
	}

	// Bind request body to Comment object
	if contentType != appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	Comment.UserID = uint(useriD)
	Comment.PhotoID = uint(photoID)

	// Create comment in database
	err = db.Debug().Create(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Return created comment
	c.JSON(http.StatusOK, gin.H{"comment": Comment})
}


func UpdateComment(c *gin.Context) {
	db := database.GetDB()
	
	contentType := helpers.GetContentType(c)

	Comment := models.Comment{}
	

	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid comment ID",
		})
		return
	}

	// Check if comment with given ID exists
	err = db.Where("id = ?", id).First(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Comment with given ID does not exist",
		})
		return
	}
	// Bind request body to Comment object
	if contentType != appJSON {
		c.ShouldBindJSON(&Comment)
	} else {
		c.ShouldBind(&Comment)
	}

	// Update comment in database
	err = db.Debug().Save(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Return updated comment
	c.JSON(http.StatusOK, gin.H{"Message": "Comment updated successfully"})
}

func DeleteComment(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid comment ID",
		})
		return
	}

	// Check if comment with given ID exists
	err = db.Where("id = ?", id).First(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Comment with given ID does not exist",
		})
		return
	}

	// Delete comment from database
	err = db.Debug().Delete(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": err.Error(),
		})
		return
	}

	// Return deleted comment
	c.JSON(http.StatusOK, gin.H{"Message": "Comment deleted successfully"})
}



func GetComments(c *gin.Context) {
	db := database.GetDB()
	Comments := []models.Comment{}

	err := db.Debug().Find(&Comments).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, Comments)
}

func GetCommentById(c *gin.Context) {
	db := database.GetDB()
	Comment := models.Comment{}

	id, err := strconv.Atoi(c.Param("commentId"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Invalid comment ID",
		})
		return
	}

	err = db.Debug().Where("id = ?", id).First(&Comment).Error
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Bad Request",
			"message": "Comment with given ID does not exist",
		})
		return
	}

	c.JSON(http.StatusOK, Comment)
}