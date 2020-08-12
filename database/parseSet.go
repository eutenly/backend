package database

import (
	"strings"
)

func parseSet(data map[string]interface{}) {

	//$set already defined
	if data["$set"] != nil {
		return
	}

	//Define $set
	data["$set"] = map[string]interface{}{}
	setData := data["$set"].(map[string]interface{})

	//Define has set
	hasSet := false

	//Loop through data
	for k, v := range data {

		//Ignore keys that start with `$`
		if strings.HasPrefix(k, "$") {
			continue
		}

		//Set data
		setData[k] = v

		//Delete key
		delete(data, k)

		//Set has set
		hasSet = true
	}

	//No values in $set
	if hasSet == false {
		delete(data, "$set")
	}
}
