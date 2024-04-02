package main

import (
	"bytes"
	"log/slog"
	"net"
	"strings"
)

// Function to handle connection and sort client options as needed
func handleConnection(conn net.Conn, logger *slog.Logger, keyString string) error {
	//Generate key from keystring
	key := []byte(keyString)

	//Log to console that connection has been made
	logger.Info(
		"Accepted connection!",
		slog.String("timestamp", generate_timestamp()),
		slog.String("ipv4_client", strings.Split(conn.RemoteAddr().String(), ":")[0]),
	)

	//Make a buffer for incoming connection
	buffer := make([]byte, 1024)

	// Read the  buffer
	_, err := conn.Read(buffer)
	if err != nil {
		logger.Error("Failed to read the initial buffer!", slog.String("timestamp", generate_timestamp()))
		return nil
	}

	//set filename - current problem is len of string is len of buffer == 1024
	cypherText := bytes.Trim(buffer, "\x00")
	if cypherText != nil {
		logger.Info(
			"Recieved cypher text!",
			slog.String("cypherText", string(cypherText)),
		)
	}

	//Get plainText and split it up
	plainTextChunk := decrypt_text(key, cypherText)
	plainText := strings.Split(plainTextChunk, ":")
	logger.Info(
		"Raw plaintext",
		slog.String("plainext", plainTextChunk),
	)

	//Verify the plainText
	if len(plainText) != 3 {
		conn.Write(encrypt_text(key, []byte("NOT ENOUGH ARGUMENTS")))
		conn.Close()
		logger.Error("Client did not provide enough arguments!", slog.String("timestamp", generate_timestamp()))
		return nil
	}

	//fmt.Println(plainText)

	//Parse plainText into options switch
	//Order is as follows
	//[0] == operation
	//[1] == suboperation
	//[2] == item
	switch plainText[0] {

	//Parse a Patch
	case "PATCH":
		logger.Info(
			"Processing client request...",
			slog.String("timestamp", generate_timestamp()),
			slog.String("ClientOperation", plainText[0]),
			slog.String("ClientSubOperation", plainText[1]),
			slog.String("ClientItem", plainText[2]),
		)

	//Parse a desired installation
	case "INSTALL":
		logger.Info(
			"Processing client request...",
			slog.String("ClientOperation", plainText[0]),
			slog.String("ClientSubOperation", plainText[1]),
			slog.String("ClientItem", plainText[2]),
		)
		//Determine desired patch
		sendBuffer := make([]byte, 1024)
		conn.Write(encrypt_text(key, []byte("SELECT")))

		//Read selection from client
		_, err = conn.Read(sendBuffer)
		if err != nil {
			return nil
		}
	}

	//make a buffer, 10GB large
	//file_buffer := make([]byte, 10000000000)

	//let client know it's ok to send a file
	//fmt.Println("Sending 'READY' signal...")
	//conn.Write([]byte("0"))

	//fmt.Println("Wrote the file!")
	return nil
}
