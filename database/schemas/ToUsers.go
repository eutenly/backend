package schemas

import (
	"go.mongodb.org/mongo-driver/bson"
)

//ToUsers converts documents to Users
func ToUsers(data []bson.M) (users []Users, returnErr error) {

	//Define result
	var result = make([]Users, len(data))

	//Loop through data
	for k, document := range data {

		//Parse data
		user, err := ToUser(document)
		if err != nil {
			returnErr = err
			return
		}

		//Add to result
		result[k] = user
	}

	//Return
	return result, nil

}
