package main

import (
	"time"

	"github.com/Azpect3120/TCPNotificationManager/internal/client"
)

func main() {
	c := client.NewTCPClient(client.WithPort(3000), client.WithTLS())
	conn := c.Configure("./certs/client.crt", "./certs/client.key", "localhost").Connect()
	defer conn.Close()

	for _, err := range c.Errors {
		panic(err)
	}

	// Simple loop to send a message to the server every 500ms
	for i := 0; i < 10; i++ {
		conn.Write([]byte("Hello, World!\n"))
		time.Sleep(500 * time.Millisecond)
	}
}
