package main

import (
	"GoGinBackend/controllers"
	"GoGinBackend/models"

	"github.com/gin-gonic/gin"
)



func main()  {
	
	r := gin.Default()

	models.ConnectDatabase()
	
	r.GET("/books", controllers.FindBooks)
	r.Run()
}