package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoadUserRoutes() *gin.Engine {
	router := gin.Default()
	router.POST("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Create user")
	})

	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, "List users")
	})

	router.GET("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Get user")
	})

	router.PUT("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Update user")
	})

	router.DELETE("/users/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Delete user")
	})

	return router
}
