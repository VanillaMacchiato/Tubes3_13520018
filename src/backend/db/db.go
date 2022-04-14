package db

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	DATABASE_NAME       = "dna_at_work"
	COLLECTION_DISEASES = "diseases"
	COLLECTION_RESULTS  = "results"
)

func Connect() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// build mongoDB connection
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("DATABASE_URL")))
	if err != nil {
		log.Fatal(err)
	}

	// ping the database
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")
	return client
}

// client instance
var DB *mongo.Client = Connect()

// collections
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database(DATABASE_NAME).Collection(collectionName)

	return collection
}
