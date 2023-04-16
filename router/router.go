package router

import ("github.com/gin-gonic/gin"
		"mygram/controllers"
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
	
	}



	return r
}