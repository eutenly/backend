package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	dbClient *mongo.Client
)

func connectDatabase() error {
	//Create Database client
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Connect to Database
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	//Check DB Connection
	fmt.Println("Checking DB Connection...")
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	dbClient = client

	collection := client.Database("eutenlyapixel").Collection("users")
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	var episodes []bson.M
	if err = cursor.All(ctx, &episodes); err != nil {
		log.Fatal(err)
	}
	fmt.Println(episodes)

	return nil
}

// func getUser(userID string) {
// 	collection := dbClient.Database("eutenlymaia").Collection("users")
// 	//collection.Find()
// }
