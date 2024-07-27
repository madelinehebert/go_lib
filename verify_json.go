package go_lib

import (
	"encoding/json"
	"log"
)

func VerifyJson(jsonData string) error {
	// Verify data converts to valid json
	_, err := json.Marshal(jsonData)
	if err != nil {
		log.Printf("%s does not look like valid JSON data.\n", jsonData)
		return err
	} else {
		return nil
	}
}
