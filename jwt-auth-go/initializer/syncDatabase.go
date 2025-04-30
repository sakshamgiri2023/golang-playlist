package initializer

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Instead of GORM's AutoMigrate, use MongoDB indexes to ensure your schema
func setupDatabase(client *mongo.Client) {
	// Get the user collection
	userCollection := client.Database("your_db_name").Collection("users")

	// Create indexes (instead of migrations)
	ctx := context.Background()
	_, err := userCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.D{{"email", 1}},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		log.Fatalf("Failed to create index: %v", err)
	}

	// Create more indexes as needed
}
