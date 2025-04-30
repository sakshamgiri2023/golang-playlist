package main

import (
	"fmt"
	"jwt-auth-go/controllers"
	"jwt-auth-go/initializer"

	"github.com/gin-gonic/gin"
)

func init() {
	initializer.LoadEnvVariable()
	initializer.ConnectToDb()
	initializer.SyncDatabase()
}
func main() {
	fmt.Println("hello world")
	router := gin.Default()
	router.GET("/Signup", controllers.Signup)
	router.Run()
}
