package main

import (
	"log"
	"os"
)

func ReadFileSize(filepath string) int {
	//Read in file's size

	/* Open a file for read access */
	if file, err := os.Open(filepath); err != nil {
		log.Println(err)
		return -1
	} else {
		defer file.Close()
		if file_stats, err := file.Stat(); err != nil {
                	log.Println(err)
			return -1
        	} else {
			return int(file_stats.Size())
		}

	}
}
