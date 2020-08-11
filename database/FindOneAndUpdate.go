package database

import (
	"context"
)

//FindOneAndUpdate updates a document with `filter` to have `data` in `collection`
func FindOneAndUpdate(collection string, filter interface{}, data map[string]interface{}) (result map[string]interface{}, err error) {

	//Parse set
	parseSet(data)

	//Run query
	err = db.Collection(collection).FindOneAndUpdate(context.Background(), filter, data).Decode(&result)
	if err != nil {
		return nil, err
	}

	//Return
	return result, nil
}
