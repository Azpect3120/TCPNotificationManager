package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Azpect3120/TCPNotificationManager/internal/client"
	"github.com/Azpect3120/TCPNotificationManager/internal/events"
	"github.com/Azpect3120/TCPNotificationManager/internal/logger"
)

func main() {
	c := client.NewTCPClient(client.WithPort(3000), client.WithTLS())
	conn := c.Configure("./certs/client.crt", "./certs/client.key", "localhost").Connect()

	// Graceful shutdown handling, capture SIGINT and SIGTERM
	// Capture Ctrl+C (SIGINT) and other termination requests (SIGTERM)
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigchan
		// Instead of using defer, this logic works the same way, for
		// testing.
		// TODO: Likely needs to change in production.
		c.Disconnect(conn)
		os.Exit(0)
	}()

	for _, err := range c.Errors {
		panic(err)
	}

	// Once connected, we need to authenticate with the server
	msg, err := json.Marshal(events.NewRequestAuthenticationEvent(""))
	if err != nil {
		panic(err)
	}

	conn.Write(msg)

	// Create a simple UI for sending messages via the terminal
	go func() {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			c.Logger.Log(fmt.Sprintf("Sending message: %s\n", line), logger.DEBUG)
			if msg, err := json.Marshal(events.NewSendMessageEvent(c.ID, line)); err != nil {
				c.Logger.Log(fmt.Sprintf("Error marshaling message: %s\n", err), logger.ERROR)
			} else {
				conn.Write(msg)
			}
		}
	}()

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
