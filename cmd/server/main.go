package main

import (
	"fmt"

	"github.com/Azpect3120/TCPNotificationManager/internal/logger"
	"github.com/Azpect3120/TCPNotificationManager/internal/server"
)

func main() {
	s := server.NewTCPServer(server.WithPort(3000), server.WithTLS(), server.WithMaxConn(2))
	ln := s.Configure("./certs/server.crt", "./certs/server.key").Listen()
	defer ln.Close()

	// Start listening
	s.Logger.Log(fmt.Sprintf("Server started on %s:%d\n", s.Opts.Addr, s.Opts.Port))

	for {
		conn, err := ln.Accept()
		if err != nil {
			s.Logger.Log(fmt.Sprintf("Error accepting connection: %s\n", err), logger.ERROR)
			return
		}
		defer conn.Close()

		go s.HandleConnection(conn)
	}
}
