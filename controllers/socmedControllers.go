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