package database

import (
	"context"
	"strings"

	"go.mongodb.org/mongo-driver/mongo"
)

//UpdateMany updates documents with `filter` to have `data` in `collection`
func UpdateMany(collection string, filter interface{}, data map[string]interface{}) (result *mongo.UpdateResult, returnErr error) {

	//Define $set
	data["$set"] = map[string]interface{}{}
	setData := data["$set"].(map[string]interface{})

	//Loop through data
	for k, v := range data {

		//Ignore keys that start with `$`
		if strings.HasPrefix(k, "$") {
			continue
		}

		//Set data
		setData[k] = v

		//Delete key
		delete(data, k)
	}

	//Run query
	result, err := db.Collection(collection).UpdateMany(context.Background(), filter, data)
	if err != nil {
		returnErr = err
		return
	}

	//Return
	return result, nil
}
