package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//FindOne finds a document with `filter` in `collection`
func FindOne(collection string, filter map[string]interface{}) (result *mongo.SingleResult) {

	//Run query
	result = db.Collection(collection).FindOne(context.Background(), filter)

	//Return
	return result
}
