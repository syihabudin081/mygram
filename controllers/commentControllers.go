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

//create comment
// @Summary Create a new comment
// @Description Create a new comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param photoId path int true "Photo ID"
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.Comment
// @Failure 400
// @Router /photos/{photoId}/comments [post]
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

// UpdateComment updates a comment
// @Summary Update a comment
// @Description Update a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param commentId path int true "Comment ID"
// @Param comment body models.Comment true "Comment"
// @Success 200 {object} models.Comment
// @Failure 400
// @Router /comments/{commentId} [put]
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

// DeleteComment deletes a comment
// @Summary Delete a comment
// @Description Delete a comment
// @Tags comments
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param commentId path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 400 
// @Router /comments/{commentId} [delete]
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


// GetComments returns all comments for a photo
// @Summary Get all comments for a photo
// @Description Get all comments for a photo
// @Tags comments
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param photoId path int true "Photo ID"
// @Success 200 {object} models.Comment
// @Failure 400 
// @Router /photos/{photoId}/comments [get]
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

// GetCommentById returns a comment by ID
// @Summary Get a comment by ID
// @Description Get a comment by ID
// @Tags comments
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param commentId path int true "Comment ID"
// @Success 200 {object} models.Comment
// @Failure 400 
// @Router /comments/{commentId} [get]
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

// GetCommentsByPhotoID returns all comments for a photo by photo ID
// @Summary Get all comments for a photo by photo ID
// @Description Get all comments for a photo by photo ID
// @Tags comments
// @Accept  json
// @Produce  json
// @Param Authorization header string true "JWT Token"
// @Param photoId path int true "Photo ID"
// @Success 200 {object} []models.Comment
// @Failure 400
// @Router /photos/{photoId}/comments [get]
func GetCommentsByPhotoID(c *gin.Context) {
    db := database.GetDB()
    Comments := []models.Comment{}

    photoID, err := strconv.Atoi(c.Param("photoID"))
	if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": "Invalid photo ID",
        })
        return
    }

    // Preload User untuk mendapatkan informasi pengguna dalam komentar
    err = db.Debug().Where("photo_id = ?", photoID).Preload("User").Find(&Comments).Error

    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{
            "error":   "Bad Request",
            "message": err.Error(),
        })
        return
    }

    c.JSON(http.StatusOK, Comments)
}

