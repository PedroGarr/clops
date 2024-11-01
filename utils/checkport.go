package utils

import (
	"fmt"
	"net"
	"time"
)

func CheckPort(host string, ports []string) {
	for _, port := range ports {
		timeout := time.Second
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		// If connection is not sucesseful
		if err != nil {
			fmt.Printf("Port %s is closed for the host %s. \n", port, host)
		}
		// If connection is sucesseful
		if conn != nil {
			fmt.Printf("Port %s is open for the host %s. \n", port, host)
			// Close the connection
			defer conn.Close()
		}
	}
}
