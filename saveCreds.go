package main

type mgoToken struct {
	iv             string `bson:"iv,omitempty"`
	ephemPublicKey string `bson:"ephemPublicKey,omitempty"`
	cipherText     string `bson:"cipherText,omitempty"`
	mac            string `bson:"mac,omitempty"`
}

type mgoConnection struct {
	ID          string   `bson:"id,omitempty"`
	accessToken mgoToken `bson:"accessToken,omitempty"`
}

type mgoUser struct {
	ID          string                   `bson:"_id,omitempty"`
	Connections map[string]mgoConnection `bson:"connections,omitempty"`
}

func storeCreds(connectionService string, userID string, accessToken string, accessSecret string) {

}
