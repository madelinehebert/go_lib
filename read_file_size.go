package go_lib

import (
	"log"
	"os"
)

/* 

   Read and return the file size of a file located at a provided file path.
   
   Negative return values are errors.
   -1 is an inability to open a file for reading.
   -2 is an inability to call os.Stat() on a file.
*/


func ReadFileSize(filepath string) int {
	/* Read in file's size */

	/* Open a file for read access */
	if file, err := os.Open(filepath); err != nil {
		log.Println(err)
		return -1
	} else {
		
		/* Defer file closing */
		defer file.Close()
		
		/* Call os.Stat() on file in order to retrieve file size. */
		if file_stats, err := file.Stat(); err != nil {
                	log.Println(err)
			return -2
        	} else {
			/* Allow defer to handle file closing and return file size to calling function. */
			return int(file_stats.Size())
		}

	}
}
