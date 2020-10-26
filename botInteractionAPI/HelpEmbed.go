package botInteractionAPI

func DispatchHelpEmbed(embed string, userID string) (err error) {
	payload := helpEmbedRequest{
		UserID: userID,
		Embed:  embed,
	}

	return requester("helpEmbed", "POST", payload)
}
