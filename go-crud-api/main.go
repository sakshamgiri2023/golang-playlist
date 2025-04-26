package main

import (
	"go-crud-api/controllers"
	"go-crud-api/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
}

func main() {

	router := gin.Default()
	router.POST("/posts", controllers.PostsCreate)
	router.Run()
}
