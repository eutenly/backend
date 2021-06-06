package schemas

//Connection is the schema for service connections
type Connection struct {
	ID           string `bson:"id,omitempty" structs:"id,omitempty"`
	AccessToken  string `bson:"accessToken,omitempty" structs:"accessToken,omitempty"`
	RefreshToken string `bson:"refreshToken,omitempty" structs:"refreshToken,omitempty"`
	AccessSecret string `bson:"accessSecret,omitempty" structs:"accessSecret,omitempty"`
	Username     string `bson:"username,omitempty" structs:"username,omitempty"`
	ConnectedAt  int    `bson:"connectedAt,omitempty" structs:"connectedAt,omitempty"`
}

//SavedLink is the schema for saved link
type SavedLink struct {
	Title       string `bson:"title,omitempty" structs:"title,omitempty"`
	Description string `bson:"description,omitempty" structs:"description,omitempty"`
	URL         string `bson:"url,omitempty" structs:"url,omitempty"`
}

//User is the schema for the users collection
type User struct {
	ID                            *string                `bson:"_id,omitempty" mapstructure:"_id" structs:"_id,omitempty"`
	OldData                       map[string]interface{} `bson:"-" structs:"oldData,omitempty"`
	Connections                   map[string]Connection  `bson:"connections,omitempty" structs:"connections,omitempty"`
	CommandsUsed                  map[string]int32       `bson:"commandsUsed,omitempty" structs:"commandsUsed,omitempty"`
	SavedLinks                    []SavedLink            `bson:"savedLinks,omitempty" structs:"savedLinks,omitempty"`
	CompactMode                   bool                   `bson:"compactMode,omitempty" structs:"compactMode,omitempty"`
	ReactionConfirmationsDisabled bool                   `bson:"reactionConfirmationsDisabled,omitempty" structs:"reactionConfirmationsDisabled,omitempty"`
	VoteExpireTimestamp           int                    `bson:"voteExpireTimestamp,omitempty" structs:"voteExpireTimestamp,omitempty"`
	PatreonTier                   int32                  `bson:"patreonTier,omitempty" structs:"patreonTier,omitempty"`
	AlphaTester                   bool                   `bson:"alphaTester,omitempty" structs:"alphaTester,omitempty"`
	BetaTester                    bool                   `bson:"betaTester,omitempty" structs:"betaTester,omitempty"`
	BetaServerOwner               bool                   `bson:"betaServerOwner,omitempty" structs:"betaServerOwner,omitempty"`
	Suggester                     bool                   `bson:"suggester,omitempty" structs:"suggester,omitempty"`
	BugHunter                     bool                   `bson:"bugHunter,omitempty" structs:"bugHunter,omitempty"`
}

////Save saves a document
//func (user User) Save() {
//
//	//Get data
//	data := structs.Map(user)
//
//	//Remove `_id` and `oldData`
//	delete(data, "_id")
//	delete(data, "oldData")
//
//	//Get changes
//	changes, _ := diff.Diff(user.OldData, data)
//
//	//Get updates
//	updates := getUpdates(changes)
//
//	//Run query
//	database.FindOneAndUpdate("users", map[string]interface{}{"_id": user.ID}, updates)
//}
//
////Delete deletes a document
//func (user User) Delete() {
//
//	//Run query
//	database.FindOneAndDelete("users", map[string]interface{}{"_id": user.ID})
//}
