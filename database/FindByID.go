package database

//FindByID finds a document by `id` in `collection`
func FindByID(collection string, id interface{}) (result map[string]interface{}, err error) {

	//Run query
	result, err = FindOne(collection, map[string]interface{}{"_id": id})
	if err != nil {
		return nil, err
	}

	//Return
	return result, nil
}
