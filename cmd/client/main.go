package main

import (
	"encoding/json"
	"errors"
	"io"
	"net"

	"github.com/Azpect3120/TCPNotificationManager/internal/client"
	"github.com/Azpect3120/TCPNotificationManager/internal/events"
)

func main() {
	c := client.NewTCPClient(client.WithPort(3000), client.WithTLS())
	conn := c.Configure("./certs/client.crt", "./certs/client.key", "localhost").Connect()
	defer conn.Close()

	for _, err := range c.Errors {
		panic(err)
	}

	// Once connected, we need to authenticate with the server
	msg, err := json.Marshal(events.NewRequestAuthenticationEvent(""))
	if err != nil {
		panic(err)
	}

	conn.Write(msg)

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if errors.Is(err, net.ErrClosed) || errors.Is(err, io.EOF) {
			// Connection closed, can exit safely
			return
		} else if err != nil {
			// Other error, for now, panic
			panic(err)
		}
		if n > 0 {
			c.HandleMessage(buf[:n])
		}
	}
}
