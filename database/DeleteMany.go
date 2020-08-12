package database

import (
	"context"
)

//DeleteMany deletes documents with `filter` in `collection`
func DeleteMany(collection string, filter interface{}) (deleted int64, returnErr error) {

	//Run query
	result, err := db.Collection(collection).DeleteMany(context.Background(), filter)
	if err != nil {
		returnErr = err
		return
	}

	//Return
	return result.DeletedCount, nil
}
