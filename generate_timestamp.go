package go_lib

import (
	"fmt"
	"time"
)

/* GenerateTimestamp is a function to generate a timestamp, mimicking the format of Google's Firestore timestamps */
func GenerateTimestamp() string {
	/* Get Date and time */
	var currentTime time.Time = time.Now()

	/* Return customized timestamp */
	return fmt.Sprintf("%s %d, %d at %d:%d:%d",
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Year(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
	)
}
