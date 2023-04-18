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


// @Summary Create a new social media account
// @Tags Social Media
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param Socmed body models.SocialMedia true "Social Media Account"
// @Success 200 {object} models.SocialMedia
// @Failure 400 
// @Router /socmeds [post]
func CreateSocmed(c *gin.Context) {
	db := database.GetDB()
	userData := c.MustGet("userData").(jwt.MapClaims)
	contentType := helpers.GetContentType(c)

	Socmed := models.SocialMedia{}
	useriD := userData["id"].(float64)

	if contentType != appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = uint(useriD)
	err := db.Debug().Create(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, Socmed)

}


// @Summary Get all social media accounts
// @Tags Social Media
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Success 200
// @Failure 400
// @Router /socmeds [get]
func GetSocialMedias(c *gin.Context) {
	db := database.GetDB()
	SocialMedias := []models.SocialMedia{}

	err := db.Debug().Find(&SocialMedias).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, SocialMedias)

}


// @Summary Get a social media account by ID
// @Tags Social Media
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param Authorization header string true "JWT Token"
// @Param socmedId path int true "Social Media ID"
// @Success 200 
// @Failure 400 
// @Router /socmeds/{socmedId} [get]
func GetSocialMediaById(c *gin.Context) {
	db := database.GetDB()
	SocialMedia := models.SocialMedia{}
	socmedId, err := strconv.Atoi(c.Param("socmedId"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	err = db.Debug().Where("id = ?", socmedId).First(&SocialMedia).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return

	}

	c.JSON(http.StatusOK, SocialMedia)

}


// UpdateSocmed updates a social media profile for the authenticated user
// @Summary Update a social media profile
// @Tags Social Media
// @Description Update a social media profile for the authenticated user
// @ID update-socmed
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param socmedId path int true "Social media profile ID"
// @Param Socmed body models.SocialMedia true "Social Media Account"
// @Success 200
// @Failure 400
// @Router /socmeds/{socmedId} [put]
func UpdateSocmed(c *gin.Context) {
	db := database.GetDB()
	contentType := helpers.GetContentType(c)

	userData := c.MustGet("userData").(jwt.MapClaims)

	Socmed := models.SocialMedia{}

	SocmedID, _ := strconv.Atoi(c.Param("socmedId"))

	userID := uint(userData["id"].(float64))

	if contentType != appJSON {
		c.ShouldBindJSON(&Socmed)
	} else {
		c.ShouldBind(&Socmed)
	}

	Socmed.UserID = userID
	Socmed.ID = uint(SocmedID)

	err := db.Debug().Where("id = ?", SocmedID).Updates(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Socmed Updated",
	})
}


// DeleteSocmed deletes a social media profile for the authenticated user
// @Summary Delete a social media profile
// @Description Delete a social media profile for the authenticated user
// @ID delete-socmed
// @Tags Social Media
// @Accept json
// @Produce json
// @Security bearerAuth
// @Param socmedId path int true "Social media profile ID"
// @Success 200 
// @Failure 400 
// @Router /socmeds/{socmedId} [delete]
func DeleteSocmed(c *gin.Context) {
	db := database.GetDB()
	Socmed := models.SocialMedia{}

	SocmedID, _ := strconv.Atoi(c.Param("socmedId"))

	err := db.Debug().Where("id = ?", SocmedID).Delete(&Socmed).Error

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"err": "Bad Request",

			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Socmed Deleted",
	})
}