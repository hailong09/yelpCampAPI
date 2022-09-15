package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB()  *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	client, _ := mongo.NewClient(options.Client().ApplyURI(uri))
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	err := client.Connect(ctx)

	if err != nil {
		log.Fatal(err)
	}

	// Ping the database
	err = client.Ping(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
	return client

}

var DB *mongo.Client = ConnectDB()

// Getting Collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("yelpcamp").Collection(collectionName)
	return collection
}