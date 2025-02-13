package main

import (
	"encoding/json"

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

	msg, err := json.Marshal(events.NewRequestAuthenticationEvent(""))
	if err != nil {
		panic(err)
	}

	conn.Write(msg)
}
