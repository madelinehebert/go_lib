package MDH

import (
	"fmt"
	"time"
)

// Function to generate a timestamp
func generate_timestamp() string {
	//Date and time
	var currentTime time.Time = time.Now()
	var timestamp string = fmt.Sprintf("%s %d, %d at %d:%d:%d",
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Year(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
	)

	//Return
	return timestamp
}
