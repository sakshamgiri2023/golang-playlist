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
	router.PUT("/posts/:id", controllers.PostsUpdate)
	router.GET("/posts", controllers.PostsIndex)
	router.GET("/posts/:id", controllers.PostsShow)
	router.DELETE("/posts/:id", controllers.PostsDelete)

	router.Run()
}
