package go_lib

import (
    "encoding/json"
    "log"
)

func VerifyJson(jsonData string) err {
    // Verify data converts to valid json
    _, json_err := json.Marshal(jsonData)
    if json_err != nil {
        log.Printf("%s does not look like valid JSON data.\n", jsonData)
	return err
    }
}
