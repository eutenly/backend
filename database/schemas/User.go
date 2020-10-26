package schemas

import (
	database ".."
	"github.com/fatih/structs"
	"github.com/r3labs/diff"
)

//Connection is the schema for service connections
type Connection struct {
	ID           string `bson:"id,omitempty" structs:"id,omitempty"`
	AccessToken  string `bson:"accessToken,omitempty" structs:"accessToken,omitempty"`
	RefreshToken string `bson:"refreshToken,omitempty" structs:"refreshToken,omitempty"`
	AccessSecret string `bson:"accessSecret,omitempty" structs:"accessSecret,omitempty"`
	ConnectedAt  int32  `bson:"connectedAt,omitempty" structs:"connectedAt,omitempty"`
}

//SavedLink is the schema for saved link
type SavedLink struct {
	URL  string `bson:"url,omitempty" structs:"url,omitempty"`
	Note string `bson:"note,omitempty" structs:"note,omitempty"`
}

//User is the schema for the users collection
type User struct {
	ID                  string                 `bson:"_id,omitempty" mapstructure:"_id" structs:"_id,omitempty"`
	OldData             map[string]interface{} `structs:"oldData,omitempty"`
	Connections         map[string]Connection  `bson:"connections,omitempty" structs:"connections,omitempty"`
	SavedLinks          []SavedLink            `bson:"savedLinks,omitempty" structs:"savedLinks,omitempty"`
	VoteExpireTimestamp int32                  `bson:"voteExpireTimestamp,omitempty" structs:"voteExpireTimestamp,omitempty"`
	PatreonTier         int32                  `bson:"patreonTier,omitempty" structs:"patreonTier,omitempty"`
	AlphaTester         bool                   `bson:"alphaTester,omitempty" structs:"alphaTester,omitempty"`
	BetaTester          bool                   `bson:"betaTester,omitempty" structs:"betaTester,omitempty"`
	BetaServerOwner     bool                   `bson:"betaServerOwner,omitempty" structs:"betaServerOwner,omitempty"`
}

//Save saves a document
func (user User) Save() {

	//Get data
	data := structs.Map(user)

	//Remove `_id` and `oldData`
	delete(data, "_id")
	delete(data, "oldData")

	//Get changes
	changes, _ := diff.Diff(user.OldData, data)

	//Get updates
	updates := getUpdates(changes)

	//Run query
	database.FindOneAndUpdate("users", map[string]interface{}{"_id": user.ID}, updates)
}

//Delete deletes a document
func (user User) Delete() {

	//Run query
	database.FindOneAndDelete("users", map[string]interface{}{"_id": user.ID})
}
