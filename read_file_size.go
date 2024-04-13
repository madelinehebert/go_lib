package go_lib

import (
	"log"
	"os"
)

func ReadFileSize(filepath string) int {
	//Read in file's size
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
	file.Close()

	return int(file_stats.Size())
}
