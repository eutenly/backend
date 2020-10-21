package authentication

import (
	"fmt"

	"../database"
	"../database/schemas"
)

func storeTokens(userID string, connectionName string, connectionUserID string, tokens map[string]interface{}) {

	//TODO: It doesnt actually modify any values and idk why
	//TODO: Upsert if user doesnt exist

	//Get user
	userRaw, _ := database.FindByID("users", userID)
	userData, _ := schemas.ToUser(userRaw)

	//Define connections
	if userData.Connections == nil {
		userData.Connections = map[string]schemas.Connection{}
	}

	//Define this connection
	if _, ok := userData.Connections[connectionName]; ok {
		userData.Connections[connectionName] = schemas.Connection{}
	}

	//Set connection data
	connectionData := userData.Connections[connectionName]
	connectionData.ID = connectionUserID
	if tokens["accessToken"] != nil {
		connectionData.AccessToken = fmt.Sprintf("%v", tokens["accessToken"])
	}
	if tokens["refreshToken"] != nil {
		connectionData.RefreshToken = fmt.Sprintf("%v", tokens["refreshToken"])
	}
	if tokens["accessSecret"] != nil {
		connectionData.AccessSecret = fmt.Sprintf("%v", tokens["accessSecret"])
	}

	//Save user
	userData.Save()
}
