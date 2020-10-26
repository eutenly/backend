package botInteractionAPI

//Schemas for API routes

//POST: /api/v1/helpEmbed
type helpEmbedRequest struct {
	UserID string `json:"user_id"`
	Embed  string `json:"embed"`
}
