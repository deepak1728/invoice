package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func ConnectToMongo() error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	userName := os.Getenv("MONGO_DB_USERNAME")
	password := os.Getenv("MONGO_DB_PASSWORD")

	clientOptions.SetAuth(options.Credential{
		Username: userName,
		Password: password,
	})

	//connection to Mongo
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	databases, err := client.ListDatabaseNames(context.Background(), bson.M{})
	if err != nil {
		fmt.Println("failed", err)
	}
	fmt.Println("databases")
	for _, dbName := range databases {
		fmt.Println(dbName)
	}
	MongoClient = client
	return nil
}

func GetCollection(collectionName string) *mongo.Collection {
	return MongoClient.Database("invoice_db").Collection(collectionName)
}
