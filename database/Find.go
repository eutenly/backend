package database

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

//Find finds documents with `filter` in `collection`
func Find(collection string, filter ...interface{}) (result []bson.M, err error) {

	//Empty filter
	if filter == nil {
		filter = []interface{}{bson.D{}}
	}

	//Run query
	cursor, err := db.Collection(collection).Find(context.Background(), filter[0])
	if err != nil {
		return nil, err
	}

	//Get all documents
	err = cursor.All(context.Background(), &result)
	if err != nil {
		return nil, err
	}

	//Return
	return result, nil
}
