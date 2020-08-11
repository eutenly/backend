package database

import (
	"context"
)

//FindOne finds a document with `filter` in `collection`
func FindOne(collection string, filter map[string]interface{}) (result map[string]interface{}, err error) {

	//Run query
	err = db.Collection(collection).FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return nil, err
	}

	//Return
	return result, nil
}
