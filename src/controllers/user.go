package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Create user")
}

func ListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, "List users")
}

func GetUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Get user")
}

func UpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Update user")
}

func DeleteUser(c *gin.Context) {
	c.JSON(http.StatusOK, "Delete user")
}
