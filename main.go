package main

import (
	"akash191095/go-crud-api/controllers"
	"akash191095/go-crud-api/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	models.ConnectDatabase()

	router.POST("/posts", controllers.CreatePost)
	router.Run()
}