package router

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()
	
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Welcome to MyGram",
		})
	},
)

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})


	userRouter := r.Group("/users")
	{
		userRouter.POST("/register",controllers.UserRegister)
		userRouter.POST("/login",controllers.UserLogin)
	
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use((middlewares.Authentication()))
		photoRouter.POST("/create",controllers.CreatePhoto)
		photoRouter.GET("/get",controllers.GetPhotos)
		photoRouter.GET("/get/:photoId",controllers.GetPhotoById)
		photoRouter.PUT("/update/:photoId",middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/delete/:photoId",middlewares.PhotoAuthorization(), controllers.DeletePhotoByID)
	}

	return r
}