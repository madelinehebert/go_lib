package go_lib

import "math/rand/v2"

func generate_random_time(min int, max int) int {
	//Return the time
	return rand.IntN(max-min) + min
}
