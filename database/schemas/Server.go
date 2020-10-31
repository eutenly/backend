package schemas

//Server is the schema for the servers collection
type Server struct {
	ID      string                 `bson:"_id,omitempty" mapstructure:"_id" structs:"_id,omitempty"`
	OldData map[string]interface{} `structs:"oldData,omitempty"`
	Prefix  string                 `bson:"prefix,omitempty" structs:"prefix,omitempty"`
}

////Save saves a document
//func (server Server) Save() {
//
//	//Get data
//	data := structs.Map(server)
//
//	//Remove `_id` and `oldData`
//	delete(data, "_id")
//	delete(data, "oldData")
//
//	//Get changes
//	changes, _ := diff.Diff(server.OldData, data)
//
//	//Get updates
//	updates := getUpdates(changes)
//
//	//Run query
//	database.FindOneAndUpdate("servers", map[string]interface{}{"_id": server.ID}, updates)
//}
//
////Delete deletes a document
//func (server Server) Delete() {
//
//	//Run query
//	database.FindOneAndDelete("servers", map[string]interface{}{"_id": server.ID})
//}
