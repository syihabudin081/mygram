package router

import (
	"mygram/controllers"
	"mygram/middlewares"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MyGram API
// @description This is a sample API for MyGram application
// @version 1
// @host localhost:8080
// @BasePath /api/v1
func StartApp() *gin.Engine {
	r := gin.Default()

	// add swagger documentation
	url := ginSwagger.URL("http://localhost:8080/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	// define endpoints
	userRouter := r.Group("/api/v1/users")
	{
		//login and register
		userRouter.POST("/register",controllers.UserRegister)
		userRouter.POST("/login",controllers.UserLogin)
	}

	photoRouter := r.Group("/api/v1/photos")
	{
		
		photoRouter.Use((middlewares.Authentication()))
		//create, get, update, delete
		photoRouter.POST("/create",controllers.CreatePhoto)
		photoRouter.GET("/get",controllers.GetPhotos)
		photoRouter.GET("/get/:photoId",controllers.GetPhotoById)
		photoRouter.PUT("/update/:photoId",middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/delete/:photoId",middlewares.PhotoAuthorization(), controllers.DeletePhotoByID)
	}

	socmedRouter := r.Group("/api/v1/socmeds")
	{

		socmedRouter.Use((middlewares.Authentication()))
		//create, get, update, delete
		socmedRouter.POST("/create",controllers.CreateSocmed)
		socmedRouter.GET("/get",controllers.GetSocialMedias)
		socmedRouter.GET("/get/:socmedId",controllers.GetSocialMediaById)
		socmedRouter.PUT("/update/:socmedId",middlewares.SocmedAuthorization(), controllers.UpdateSocmed)
		socmedRouter.DELETE("/delete/:socmedId",middlewares.SocmedAuthorization(), controllers.DeleteSocmed)
	}

	commentRouter := r.Group("/api/v1/comments")
	{
		commentRouter.Use((middlewares.Authentication()))
		//create, get, update, delete
		commentRouter.POST("/create/:photoId", controllers.CreateComment)
		commentRouter.GET("/get/", controllers.GetComments)
		commentRouter.GET("/get/:commentId", controllers.GetCommentById)
		commentRouter.PUT("/update/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/delete/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	return r
}