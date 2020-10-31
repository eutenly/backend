package botInteractionAPI

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func requester(route string, method string, payload interface{}) (err error) {
	//Construct body of request
	query, err := json.Marshal(payload)
	if err != nil {
		return
	}

	//Make request
	endpoint := fmt.Sprintf("%v/api/v1/%v", os.Getenv("BOT_ENDPOINT"), route)
	request, err := http.NewRequest(method, endpoint, bytes.NewBuffer(query))
	if err != nil {
		return
	}
	request.Header.Set("User-Agent", "eutenly-backend/0.1")
	request.Header.Set("Authorization", os.Getenv("BOT_ENDPOINT_KEY"))

	//Send request
	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	return
}
