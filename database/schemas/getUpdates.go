package schemas

import (
	"strings"

	"github.com/r3labs/diff"
)

func getUpdates(changes diff.Changelog) (updates map[string]interface{}) {

	//Define updates
	updates = map[string]interface{}{}

	//Loop through changes
	for _, change := range changes {

		//Create and update
		if change.Type == "create" || change.Type == "update" {

			//Define $set
			if updates["$set"] == nil {
				updates["$set"] = map[string]interface{}{}
			}

			setData := updates["$set"].(map[string]interface{})

			//Set data
			setData[strings.Join(change.Path, ".")] = change.To
		}

		//Delete
		if change.Type == "delete" {

			//Define $unset
			if updates["$unset"] == nil {
				updates["$unset"] = map[string]interface{}{}
			}

			unsetData := updates["$unset"].(map[string]interface{})

			//Set data
			unsetData[strings.Join(change.Path, ".")] = 1
		}
	}

	//Return
	return updates
}
