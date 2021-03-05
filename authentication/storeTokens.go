package authentication

import (
	"eutenly/backend/database"
	"eutenly/backend/database/schemas"

	"github.com/sirupsen/logrus"
	"time"
)

func storeTokens(userID string, connectionName string, connectionUserID string, connectionUsername string, tokens map[string]string) (err error) {
	// Get user
	user, err := database.GetUser(userID)
	if err != nil {
		logrus.Error(err)
		return
	}

	// Create connection struct
	connection := schemas.Connection{
		ID:           connectionUserID,
		Username:     connectionUsername,
		AccessToken:  tokens["accessToken"],
		RefreshToken: tokens["refreshToken"],
		AccessSecret: tokens["accessSecret"],
		ConnectedAt:  int(time.Now().UnixNano() / int64(time.Millisecond)),
	}

	// Insert connection to user
	if user.Connections == nil {
		user.Connections = make(map[string]schemas.Connection)
	}
	user.Connections[connectionName] = connection

	// Store user
	err = database.SetUser(user)
	if err != nil {
		logrus.Error(err)
		return
	}

	//Stats
	//influxdb.CollectStat("accounts_authorized", nil, map[string]interface{}{"type": connectionName})

	return
}
