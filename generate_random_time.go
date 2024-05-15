package go_lib

import "math/rand/v2"

func GenerateRandomTime(min int, max int) int {
	//Return the time
	return rand.IntN(max-min) + min
}
