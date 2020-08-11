package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

//UpdateMany updates documents with `filter` to have `data` in `collection`
func UpdateMany(collection string, filter interface{}, data map[string]interface{}) (result *mongo.UpdateResult, returnErr error) {

	//Parse set
	parseSet(data)

	//Run query
	result, err := db.Collection(collection).UpdateMany(context.Background(), filter, data)
	if err != nil {
		returnErr = err
		return
	}

	//Return
	return result, nil
}
