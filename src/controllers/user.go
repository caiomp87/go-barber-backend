package controllers

import (
	"barber/src/models"
	"barber/src/repositories"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var user *models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to read data from request: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = repositories.UserCollection.Create(ctx, user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "unable to create an user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "user created successfully",
	})
}

func ListUsers(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// tries to fetch bots list from database
	users, err := repositories.UserCollection.List(ctx)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "unable to list users",
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user, err := repositories.UserCollection.FindByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "unable to find an user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user *models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable to read data from request: " + err.Error(),
		})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = repositories.UserCollection.UpdateByID(ctx, id, user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable update an user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user updated successfully",
	})
}

func DeleteUser(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := repositories.UserCollection.DeleteByID(ctx, id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "unable remove an user: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "user removed successfully",
	})
}
