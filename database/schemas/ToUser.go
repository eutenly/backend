package schemas

import (
	"github.com/mitchellh/mapstructure"
)

//ToUser converts a document to a User
func ToUser(data map[string]interface{}) (user Users, returnErr error) {

	//Parse data
	err := mapstructure.Decode(data, &user)
	if err != nil {
		returnErr = err
		return
	}

	//Return
	return user, nil

}
