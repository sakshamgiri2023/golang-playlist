package controllers

import (
	"context"
	"go-crud-api/initializers"
	"go-crud-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func PostsCreate(c *gin.Context) {
	// get data off req body
	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	//create the post
	post := models.Post{Title: body.Title, Body: body.Body}

	collection := initializers.MongoDB.Collection("posts")
	result, err := collection.InsertOne(context.TODO(), post)
	if err != nil {
		c.JSON(500, gin.H{"error": "failed to create post"})
		return
	}

	//return it
	c.JSON(200, gin.H{
		"posts": result.InsertedID,
	})
}

func PostsIndex(c *gin.Context) {
	collection := initializers.MongoDB.Collection("posts")
	cursor, err := collection.Find(context.TODO(), bson.M{})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to fetch posts"})
		return
	}
	defer cursor.Close(context.TODO())

	var posts []models.Post
	if err := cursor.All(context.TODO(), &posts); err != nil {
		c.JSON(500, gin.H{"error": "Failed to parse posts"})
		return
	}

	c.JSON(200, gin.H{"posts": posts})
}

func PostsShow(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := initializers.MongoDB.Collection("posts")
	var post models.Post
	err = collection.FindOne(context.TODO(), bson.M{"_id": objectID}).Decode(&post)
	if err != nil {
		c.JSON(404, gin.H{"error": "Post not found"})
		return
	}

	c.JSON(200, gin.H{"post": post})
}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	var body struct {
		Body  string `json:"body"`
		Title string `json:"title"`
	}
	if err := c.Bind(&body); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	collection := initializers.MongoDB.Collection("posts")
	update := bson.M{"$set": bson.M{"title": body.Title, "body": body.Body}}
	_, err = collection.UpdateOne(context.TODO(), bson.M{"_id": objectID}, update)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to update post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post updated successfully"})
}

func PostsDelete(c *gin.Context) {
	id := c.Param("id")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID format"})
		return
	}

	collection := initializers.MongoDB.Collection("posts")
	_, err = collection.DeleteOne(context.TODO(), bson.M{"_id": objectID})
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to delete post"})
		return
	}

	c.JSON(200, gin.H{"message": "Post deleted successfully"})
}
