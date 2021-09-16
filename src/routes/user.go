package routes

import (
	"barber/src/controllers"

	"github.com/gin-gonic/gin"
)

func LoadUserRoutes() *gin.Engine {
	router := gin.Default()

	router.POST("/users", controllers.CreateUser)
	router.GET("/users", controllers.ListUsers)
	router.GET("/users/:id", controllers.GetUser)
	router.PUT("/users/:id", controllers.UpdateUser)
	router.DELETE("/users/:id", controllers.DeleteUser)

	return router
}
