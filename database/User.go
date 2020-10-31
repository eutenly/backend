package database

import (
	"./schemas"
	"context"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetUser(UserID string) (user schemas.User, err error) {
	// Fetch user
	query := db.Collection("users").FindOne(context.TODO(), bson.M{"_id": UserID}).Decode(&user)

	// If there's not documents, return the default user
	if query == mongo.ErrNoDocuments {
		// Return default values
		return schemas.User{ID: &UserID}, nil
	}

	// Check for other errors
	if query != nil {
		logrus.Error(query)
		return
	}

	// Returned decoded user
	return
}

func SetUser(user schemas.User) (err error) {
	// Create filter
	match := bson.M{"_id": user.ID}

	// Remove ID from struct
	user.ID = nil

	// Store user
	_, err = db.Collection("users").UpdateOne(context.TODO(), match, bson.M{"$set": user}, options.Update().SetUpsert(true))

	// Check for errors
	if err != nil {
		logrus.Error(err)
		return
	}

	// Return
	return
}

func UnsetUserConnection(UserID string) (err error) {
	// Create filter
	match := bson.M{"_id": UserID}

	// Store user
	_, err = db.Collection("users").UpdateOne(context.TODO(), match, bson.M{"$unset": bson.M{"connections": ""}}, options.Update().SetUpsert(true))

	// Check for errors
	if err != nil {
		logrus.Error(err)
		return
	}

	// Return
	return
}
