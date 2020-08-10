package database

import (
	"go.mongodb.org/mongo-driver/mongo"
)

//FindByID finds a document by `id` in `collection`
func FindByID(collection string, id interface{}) (result *mongo.SingleResult) {

	//Run query
	result = FindOne(collection, map[string]interface{}{"_id": id})

	//Return
	return result
}
