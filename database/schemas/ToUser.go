package schemas

import (
	"go.mongodb.org/mongo-driver/mongo"
)

//ToUser converts a document to a User
func ToUser(data *mongo.SingleResult) (user Users, returnErr error) {

	err := data.Decode(&user)
	if err != nil {
		returnErr = err
		return
	}

	//Return
	return user, nil

}
