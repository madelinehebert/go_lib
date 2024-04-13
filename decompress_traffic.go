package go_lib

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"io"
	"log"
)

// Function to decompress traffic
func DecompressTraffic(conn tls.Conn) []byte {
	//Make a buffer
	var ibuf bytes.Buffer

	//Make a new gzip reader
	zr, err := gzip.NewReader(&conn)
	if err != nil {
		log.Fatal(err)
	}
	zr.Close()

	//Write data to file
	_, err = io.Copy(&ibuf, zr)
	if err != nil {
		log.Fatal(err)
	}

	return ibuf.Bytes()
}
