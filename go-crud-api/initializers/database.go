package initializers

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client
var MongoDB *mongo.Database

func ConnectToDb() {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable.")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	// Ping the database to verify connection
	if err := client.Ping(context.TODO(), nil); err != nil {
		log.Fatal("Failed to ping MongoDB:", err)
	}

	// Set the global client and database
	MongoClient = client
	MongoDB = client.Database("your_database_name") // Replace with your database name

	log.Println("Connected to MongoDB successfully")
}
