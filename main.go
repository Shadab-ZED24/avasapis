package main

import (
	"github.com/fygni/avasapis/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		userController := new(controllers.UserController)
		v1.POST("/users", userController.Create)
		v1.POST("/login", userController.Login)
		v1.POST("/addgroup", userController.AddConf)
		v1.GET("/users", userController.GetAll)
		v1.GET("/users/:id", userController.GetOne)
		v1.DELETE("/users/:id", userController.DeleteOne)
		v1.PUT("/users/:id", userController.UpdateOne)
		v1.PUT("/updategroup/:id", userController.UpdateGroup)
		v1.PUT("/deletegroup/:id", userController.DeleteGroup)
	}

	router.NoRoute(func(context *gin.Context) {
		context.String(404, "fuck you")
	})

	router.Run()
}
