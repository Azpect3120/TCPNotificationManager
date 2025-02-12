package main

import (
	"fmt"

	"github.com/Azpect3120/TCPNotificationManager/internal/server"
)

func main() {
	s := server.NewTCPServer(server.WithPort(3000), server.WithTLS())
	ln := s.Configure("./certs/server.crt", "./certs/server.key").Listen()
	fmt.Printf("Server: %v\n", s)
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err)
		}
		defer conn.Close()

		go s.HandleConnection(conn)

	}
}
