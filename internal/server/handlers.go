package server

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/Azpect3120/TCPNotificationManager/internal/events"
	"github.com/Azpect3120/TCPNotificationManager/internal/utils"
)

// When the client sends a request to authenticate, this function will be called.
// This function will handle the request and send a response back to the client.
//
// This function assumes there is space in the server for the client to connect,
// as it was already confirmed that there is. This function also assumes that
// the client is not already authenticated but exists in the Conns slice in
// the server. If it is not found, an error will be thrown.
func RequestAuthenticationHandler(server *TcpServer, conn net.Conn, event *events.RequestAuthenticationEvent) {
	var exists bool = false
	for _, c := range server.Conns {
		if c == conn {
			exists = true
			break
		}
	}
	if !exists {
		// Send back a rejected message
		fmt.Printf("Client not found in server connections: %s\n", conn.RemoteAddr().String())
		return
	}

	// Authenticate the client
	clientId := utils.GenerateClientID()
	server.Authorized[clientId] = conn

	// Display a message for now, but in the future, this can be an event
	// to all other client, that a new client has been accepted.
	fmt.Printf("A client '%s' has been authenticated\n", clientId)
	for _, c := range server.Conns {
		fmt.Printf("Client: %s\n", c.RemoteAddr().String())
	}
	fmt.Printf("%+v\n", server.Authorized)

	// Send back the message to the client
	response := events.NewConnectionAcceptedEvent(server.ID, clientId)
	if bytes, err := json.Marshal(response); err != nil {
		fmt.Printf("Error marshalling response: %s\n", err)
	} else {
		conn.Write(bytes)
	}
}
