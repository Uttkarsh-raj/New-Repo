package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	// "github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client = DBInstance()

func DBInstance() *mongo.Client {
	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading the env to set the db: ", err.Error())
	// }
	MongoDb := os.Getenv("MONGODB_URL")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MongoDb))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB.")
	return client
}

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("cluster0").Collection(collectionName)
}
