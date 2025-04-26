package controllers

import (
	"go-crud-api/initializers"
	"go-crud-api/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	//create the post
	post := models.Post{Title: body.Title, Body: body.Body}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	//return it
	c.JSON(200, gin.H{
		"post": post,
	})
}
