package go_lib

import (
	"bytes"
	//"compress/gzip"
	"crypto/tls"
	"io"
	"log"

	"github.com/klauspost/compress/zstd"
)

// Function to decompress traffic
func DecompressTraffic(conn tls.Conn) []byte {
	//Make a buffer
	var ibuf bytes.Buffer

	//Make a new gzip reader
	//zr, err := gzip.NewReader(&conn)
	zr, err := zstd.NewReader(&conn)
	if err != nil {
		log.Fatal(err)
	}

	//Write data to file
	_, err = io.Copy(&ibuf, zr)
	if err != nil {
		log.Fatal(err)
	}

	zr.Close()

	return ibuf.Bytes()
}
