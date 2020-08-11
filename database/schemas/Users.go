package schemas

//Token is the schema for tokens
type Token struct {
	IV             string `bson:"iv,omitempty"`
	EphemPublicKey string `bson:"ephemPublicKey,omitempty"`
	CipherText     string `bson:"cipherText,omitempty"`
	Mac            string `bson:"mac,omitempty"`
}

//Connection is the schema for service connections
type Connection struct {
	ID           string `bson:"id,omitempty"`
	AccessToken  Token  `bson:"accessToken,omitempty"`
	RefreshToken Token  `bson:"refreshToken,omitempty"`
	AccessSecret Token  `bson:"accessSecret,omitempty"`
}

//SavedLink is the schema for saved link
type SavedLink struct {
	URL  string `bson:"url,omitempty"`
	Note string `bson:"note,omitempty"`
}

//Users is the schema for the users collection
type Users struct {
	ID                  string                `bson:"_id,omitempty" mapstructure:"_id"`
	Connections         map[string]Connection `bson:"connections,omitempty"`
	SavedLinks          []SavedLink           `bson:"savedLinks,omitempty"`
	VoteExpireTimestamp int32                 `bson:"voteExpireTimestamp,omitempty"`
	PatreonTier         int32                 `bson:"patreonTier,omitempty"`
}
