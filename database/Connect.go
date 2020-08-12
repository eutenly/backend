package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	db *mongo.Database
)

//Connect connects to the database using the `MONGO_URI` and `MONGO_DATABASE` env variables
func Connect() error {

	//Create database client
	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	//Connect to database
	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	//Check database connection
	fmt.Println("Checking database connection...")

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	fmt.Println("Database is connected")

	//Get db
	db = client.Database(os.Getenv("MONGO_DATABASE"))

	//Return
	return nil
}
