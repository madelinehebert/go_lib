package go_lib

import (
	"fmt"
	"log"
	"net"
)

/* Get the outbound ip of this machine. */
func GetLocalIP() string {

	/* Dial Google's public DNS server. */
	if conn, err := net.Dial("udp", "8.8.8.8:80"); err != nil {
		log.Println(err)
		return "BADDIAL"
	} else {
		/* Defer closing the connection. */
		defer conn.Close()

		/* Return outbound IP address. */
		return fmt.Sprint(conn.LocalAddr().(*net.UDPAddr).IP)
	}

}
