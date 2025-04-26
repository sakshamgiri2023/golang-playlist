package main

import (
	"go-crud-api/initializers"
	"go-crud-api/models"
)

func init() {
	initializers.LoadEnvVariable()
	initializers.ConnectToDb()
}

func main() {
	initializers.DB.AutoMigrate(&models.Post{})
}
