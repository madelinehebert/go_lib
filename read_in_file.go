package main

import (
	"log"
	"os"
)

// Function to read in a file
func ReadInFile(filepath string) ([]byte, error) {
	//Read in file
	/* Open a file for read access */
	file, err := os.Open(filepath)
	if err != nil {
		log.Fatal(err)
	}

	/* find the file on the system, find the amount of bytes it occupies */
	file_stats, err := file.Stat()
	if err != nil {
		log.Println(err)
	}

	/* Make an array of bytes, with a maximum length of 100 */
	data := make([]byte, file_stats.Size())

	/* Data in read from the file, storing the number of bytes and any errors */
	file.Read(data)

	file.Close()
	//Return the file bytes
	return data, nil
}
