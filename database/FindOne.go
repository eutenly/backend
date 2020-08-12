package database

import (
	"context"
)

//FindOne finds a document with `filter` in `collection`
func FindOne(collection string, filter ...map[string]interface{}) (result map[string]interface{}, err error) {

	//Empty filter
	if filter == nil {
		filter = []map[string]interface{}{{}}
	}

	//Run query
	err = db.Collection(collection).FindOne(context.Background(), filter[0]).Decode(&result)
	if err != nil {
		return nil, err
	}

	//Return
	return result, nil
}
