package database

import (
	"context"
)

//InsertOne inserts a document with `data` into `collection`
func InsertOne(collection string, data interface{}) (id interface{}, err error) {

	//Run query
	result, err := db.Collection(collection).InsertOne(context.Background(), data)
	if err != nil {
		return nil, err
	}

	//Return
	return result.InsertedID, nil
}
