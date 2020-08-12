package schemas

import (
	"github.com/mitchellh/mapstructure"
)

//ToUser converts a document to a User
func ToUser(data map[string]interface{}) (user User, returnErr error) {

	//Parse data
	err := mapstructure.Decode(data, &user)
	if err != nil {
		returnErr = err
		return
	}

	//Set old data
	delete(data, "_id")
	user.OldData = data

	//Return
	return user, nil

}
