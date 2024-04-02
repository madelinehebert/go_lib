package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	toml "github.com/pelletier/go-toml"
)

// Define constants
const IP = "127.0.0.1"
const PORT = ":63455"
const LOG_FILE = "log/MyLogFile.txt"
const CERT_LOCATION = "etc/cert.pem"
const CONFIG_LOCATION = "etc/UpdateServerConfig.toml"
const KEY_LOCATION = "etc/private_key.pem"
const KEY_STRING = "passphrasewhichneedstobe32bytes!"

// Config struct
type UpdateServerConfig struct {
	PORT          string
	IP            string
	KEY_STRING    string
	KEY_LOCATION  string
	CERT_LOCATION string
	LOG_FILE      string
}

func main() {
	//Read in config
	cfg := read_in_config()

	//Open log file, attached writer
	f, err := os.OpenFile(LOG_FILE, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	//Prepare to catch SIGINT
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\nCaught SIGINT!\nExiting!")
		w.Flush()
		os.Exit(1)
	}()

	//Create custom logger, set as default
	logger := slog.New(slog.NewJSONHandler(w, nil))
	slog.SetDefault(logger)

	//Log start up
	logger.Info(
		"Starting server...",
		slog.String("ipv4", cfg.IP),
		slog.String("port", PORT[1:]),
		slog.Int("pid", os.Getpid()),
		slog.String("user", os.Getenv("USER")),
		slog.String("timestamp", generate_timestamp()),
	)

	//Set log flags
	log.SetFlags(log.Lshortfile)

	//Import cert and key
	cer, err := tls.LoadX509KeyPair(CERT_LOCATION, KEY_LOCATION)
	if err != nil {
		log.Println(err)
		return
	}

	//Configure tls
	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	//Start TLS listener
	ln, err := tls.Listen("tcp", PORT, config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	//Handle connections
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConnection(conn, logger)
	}
}

// Function to read in config
func read_in_config() UpdateServerConfig {
	// Reading from a TOML file
	/*
		_, err := toml.LoadFile(CONFIG_LOCATION)
		if err != nil {
			fmt.Println("Error reading TOML file:", err)
			return
		}
		// Accessing values from the TOML tree
		//t := data.Get("system").(*toml.Tree)
		//fmt.Println(t)
		//fmt.Println(data.Get("system.PORT"))
		// Access individual values
		//service_name := data.Get("svcbundle.service_name")
		//fmt.Println("service_name : ", service_name)
	*/

	//Auto unmarshall config to struct
	cfg := UpdateServerConfig{}
	myBytes, _ := os.ReadFile(CONFIG_LOCATION)

	toml.Unmarshal(myBytes, &cfg)
	return cfg

}

// Function to read in a file
func read_in_file(filename string) []byte {
	//Read in file data
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	//Return file data
	return data
}

func send_binary(conn tls.Conn, filename string) {
	//Retrieve file data
	data := read_in_file(filename)
	fmt.Println(data)
}

// Function to generate a timestamp
func generate_timestamp() string {
	//Date and time
	var currentTime time.Time = time.Now()
	var timestamp string = fmt.Sprintf("%s %d, %d at %d:%d:%d",
		currentTime.Month(),
		currentTime.Day(),
		currentTime.Year(),
		currentTime.Hour(),
		currentTime.Minute(),
		currentTime.Second(),
	)

	//Return
	return timestamp
}
