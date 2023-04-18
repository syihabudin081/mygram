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

	socmedRouter := r.Group("/socmeds")
	{
	socmedRouter.Use((middlewares.Authentication()))
	socmedRouter.POST("/create",controllers.CreateSocmed)
	socmedRouter.GET("/get",controllers.GetSocialMedias)
	socmedRouter.GET("/get/:socmedId",controllers.GetSocialMediaById)
	socmedRouter.PUT("/update/:socmedId",middlewares.SocmedAuthorization(), controllers.UpdateSocmed)
	socmedRouter.DELETE("/delete/:socmedId",middlewares.SocmedAuthorization(), controllers.DeleteSocmed)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use((middlewares.Authentication()))
		commentRouter.POST("/create/:photoId", controllers.CreateComment)
		commentRouter.GET("/get/", controllers.GetComments)
		commentRouter.GET("/get/:commentId", controllers.GetCommentById)
		commentRouter.PUT("/update/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/delete/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	

	return r
}