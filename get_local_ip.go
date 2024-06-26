package go_lib

import (
	"fmt"
	"log"
	"net"
)

// Get preferred outbound ip of this machine
func GetLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return fmt.Sprint(localAddr.IP)
}
