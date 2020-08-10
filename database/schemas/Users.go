package schemas

//Users is the schema for the users collection
type Users struct {
	ID                  string `bson:"_id"`
	VoteExpireTimestamp int    `bson:"voteExpireTimestamp"`
	PatreonTier         int    `bson:"patreonTier"`
}
