package schemas

import (
	"github.com/mitchellh/mapstructure"
)

//ToServer converts a document to a server
func ToServer(data map[string]interface{}) (server Server, returnErr error) {

	//Parse data
	err := mapstructure.Decode(data, &server)
	if err != nil {
		returnErr = err
		return
	}

	//Set old data
	delete(data, "_id")
	server.OldData = data

	//Return
	return server, nil

}
