package database

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo/options"
)

//FindOneAndUpdate updates a document with `filter` to have `data` in `collection`
func FindOneAndUpdate(collection string, filter interface{}, data map[string]interface{}, upsert ...bool) (result map[string]interface{}, err error) {

	//Parse set
	parseSet(data)

	//Parse options
	if upsert == nil {
		upsert = []bool{false}
	}
	options := options.FindOneAndUpdate().SetUpsert(upsert[0])

	//Run query
	err = db.Collection(collection).FindOneAndUpdate(context.Background(), filter, data, options).Decode(&result)
	if err != nil {
		return nil, err
	}

	//Return
	return result, nil
}
