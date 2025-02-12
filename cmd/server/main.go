package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/Azpect3120/TCPNotificationManager/internal/server"
)

func handleClient(conn net.Conn) {
	fmt.Println(conn.RemoteAddr().String())
	var buf [1024]byte

	for {
		n, err := conn.Read(buf[:])
		if err != nil && !errors.Is(err, net.ErrClosed) && !errors.Is(err, io.EOF) {
			log.Fatalln(err)
		}
		if n > 0 {
			fmt.Print("Received: ", string(buf[:n]))
		}
	}
}

func main() {

	// listener, err := net.Listen("tcp", ":8080")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer listener.Close()
	//
	// for {
	// 	conn, err := listener.Accept()
	// 	if err != nil {
	// 		log.Fatalln(err)
	// 		continue
	// 	}
	// 	defer conn.Close()
	//
	// 	go handleClient(conn)
	// }

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

		go handleClient(conn)

	}
}
