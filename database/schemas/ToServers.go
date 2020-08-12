package schemas

import (
	"go.mongodb.org/mongo-driver/bson"
)

//ToServers converts documents to Servers
func ToServers(data []bson.M) (servers []Servers, returnErr error) {

	//Define result
	var result = make([]Servers, len(data))

	//Loop through data
	for k, document := range data {

		//Parse data
		server, err := ToServer(document)
		if err != nil {
			returnErr = err
			return
		}

		//Add to result
		result[k] = server
	}

	//Return
	return result, nil

}
