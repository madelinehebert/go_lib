package go_lib

import (
	"bufio"
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Compress network traffic
func CompressTraffic(filepath string) []byte {
	//Open the file for sending
	file, err := os.Open(strings.TrimSpace(filepath)) // For read access.
	if err != nil {
		fmt.Println(err.Error())
		return []byte("")
	}
	defer file.Close()

	//Set up a buffer
	var obuf bytes.Buffer

	//Set up a reader
	r := bufio.NewReader(file)

	//Compress the input data
	zw := gzip.NewWriter(&obuf)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	//Copy over data and close gz writer
	_, err = io.Copy(zw, r)
	if err != nil {
		log.Fatal(err)
	}
	zw.Flush()
	zw.Close()

	return obuf.Bytes()
}
