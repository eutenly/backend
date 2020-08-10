package schemas

type token struct {
	IV             string `bson:"iv,omitempty"`
	EphemPublicKey string `bson:"ephemPublicKey,omitempty"`
	CipherText     string `bson:"cipherText,omitempty"`
	Mac            string `bson:"mac,omitempty"`
}

type connection struct {
	ID           string `bson:"id,omitempty"`
	AccessToken  token  `bson:"accessToken,omitempty"`
	RefreshToken token  `bson:"refreshToken,omitempty"`
	AccessSecret token  `bson:"accessSecret,omitempty"`
}

type savedLink struct {
	URL  string `bson:"url,omitempty"`
	Note string `bson:"note,omitempty"`
}

//Users is the schema for the users collection
type Users struct {
	ID                  string                `bson:"_id,omitempty"`
	Connections         map[string]connection `bson:"connections,omitempty"`
	SavedLinks          []savedLink           `bson:"savedLinks,omitempty"`
	VoteExpireTimestamp int32                 `bson:"voteExpireTimestamp,omitempty"`
	PatreonTier         int32                 `bson:"patreonTier,omitempty"`
}
