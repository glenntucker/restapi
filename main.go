package main

import (
	"githib.com/glenntucker/restapi/controllers"
	"githib.com/glenntucker/restapi/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDatabase()

	r.GET("/books", controllers.FindBooks) // new

	r.Run()
}
