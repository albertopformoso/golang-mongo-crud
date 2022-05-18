package config

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB(collectionName string) (collection *mongo.Collection) {

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))

	// Connect to Mongo
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")

	collection = client.Database(os.Getenv("DATABASE")).Collection(collectionName)

	return
}
