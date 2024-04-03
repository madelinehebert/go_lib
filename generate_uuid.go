package go_lib

import (
	"log"
	"math/rand"
	"strconv"
)

// A constant to hold all the valid characters to be used in uuid creation
const letterBytes = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// Function to generate a unique uuid for each document
func GenerateUUID(n string) string {

	//Convert n to integer
	n_int, err := strconv.Atoi(n)
	if err != nil {
		log.Println(err)
		return "BADINT"
	}

	//Make a byte slice of n length
	uuid := make([]byte, n_int)

	//Create uuid randomly
	for i := range uuid {
		uuid[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	//return uuid
	//log.Println(string(uuid))
	return string(uuid)
}
