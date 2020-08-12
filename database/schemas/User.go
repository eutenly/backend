package schemas

import (
	database ".."
	"github.com/fatih/structs"
	"github.com/r3labs/diff"
)

//Token is the schema for tokens
type Token struct {
	IV             string `bson:"iv,omitempty" structs:"iv,omitempty"`
	EphemPublicKey string `bson:"ephemPublicKey,omitempty" structs:"ephemPublicKey,omitempty"`
	CipherText     string `bson:"cipherText,omitempty" structs:"cipherText,omitempty"`
	Mac            string `bson:"mac,omitempty" structs:"mac,omitempty"`
}

//Connection is the schema for service connections
type Connection struct {
	ID           string `bson:"id,omitempty" structs:"id,omitempty"`
	AccessToken  Token  `bson:"accessToken,omitempty" structs:"accessToken,omitempty"`
	RefreshToken Token  `bson:"refreshToken,omitempty" structs:"refreshToken,omitempty"`
	AccessSecret Token  `bson:"accessSecret,omitempty" structs:"accessSecret,omitempty"`
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
