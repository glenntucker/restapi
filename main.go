package main

import (
	"github.com/gin-gonic/gin"
	"github.com/glenntucker/restapi/controllers"
	"github.com/glenntucker/restapi/models"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
