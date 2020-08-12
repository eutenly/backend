package database

import (
	"context"
)

//InsertMany inserts documents with `data` into `collection`
func InsertMany(collection string, data []interface{}) (id []interface{}, err error) {

	//Run query
	result, err := db.Collection(collection).InsertMany(context.Background(), data)
	if err != nil {
		return nil, err
	}

	//Return
	return result.InsertedIDs, nil
}
