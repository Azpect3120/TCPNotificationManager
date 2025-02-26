package server

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/Azpect3120/TCPNotificationManager/internal/events"
	"github.com/Azpect3120/TCPNotificationManager/internal/logger"
	"github.com/Azpect3120/TCPNotificationManager/internal/utils"
)

// RequestAuthenticationHandler When the client sends a request to authenticate,
// this function will be called.
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
		server.Logger.Log(fmt.Sprintf("Client not found in server connections: %s\n", conn.RemoteAddr().String()), logger.ERROR)
		return
	}

	// Authenticate the client
	clientId := utils.GenerateClientID()
	server.Authorized[clientId] = conn

	// Display a message for now, but in the future, this can be an event
	// to all other client, that a new client has been accepted.
	server.Logger.Log(fmt.Sprintf("A client '%s' has been authenticated\n", clientId))

	// Send back the message to the client
	if bytes, err := json.Marshal(events.NewConnectionAcceptedEvent(server.ID, clientId)); err != nil {
		server.Logger.Log(fmt.Sprintf("Error marshalling response: %s\n", err), logger.ERROR)
	} else {
		conn.Write(bytes)
	}

	// Client has been authenticated, now we can broadcast the message to all clients
	message, err := json.Marshal(events.NewClientAuthenticatedEvent(server.ID, clientId))
	if err != nil {
		server.Logger.Log(fmt.Sprintf("Error marshalling response: %s\n", err), logger.ERROR)
	} else {
		errs := server.BroadcastMessage(message, conn)
		for _, err := range errs {
			server.Logger.Log(fmt.Sprintf("Error broadcasting message: %s\n", err), logger.ERROR)
		}
	}
}

// ClientDisconnectingHandler When a client disconnects from the server, this function
// will be called on the server.
// This function will handle the disconnection and remove the client from the authenticated map.
// Additionally, the server will broadcast the disconnection event to all other clients.
//
// Each client is removed from the server's connections slice when they disconnect,
// so there is no need to remove them here.
func ClientDisconnectingHandler(server *TcpServer, conn net.Conn, event *events.ClientDisconnectingEvent) {
	// Delete from the authorized map
	delete(server.Authorized, event.ID)

	server.Logger.Log(fmt.Sprintf("Client '%s' has disconnected\n", event.ID), logger.DEBUG)

	// TODO: Broadcast the message to all clients
	message, err := json.Marshal(events.NewClientDisconnectedEvent(server.ID, event.ID))
	if err != nil {
		server.Logger.Log(fmt.Sprintf("Error marshalling response: %s\n", err), logger.ERROR)
	} else {
		errs := server.BroadcastMessage(message, conn)
		for _, err := range errs {
			server.Logger.Log(fmt.Sprintf("Error broadcasting message: %s\n", err), logger.ERROR)
		}
	}
}

// SendMessageHandler When a client sends a message to the server, this function will be called.
// This function will handle the message and broadcast it to all other clients.
//
// Handling the message will include checking if the client is authenticated, and
// if the message is valid. If the client is not authenticated, the message will
// be ignored.
func SendMessageHandler(server *TcpServer, conn net.Conn, event *events.SendMessageEvent) {
	// Check if the client is authenticated
	if event.ID == "" {
		server.Logger.Log(fmt.Sprintf("Client ID is empty\n"), logger.ERROR)
	} else if _, ok := server.Authorized[event.ID]; !ok {
		server.Logger.Log(fmt.Sprintf("Client '%s' is not authenticated\n", event.ID), logger.ERROR)
		return
	}

	// Broadcast the message to all clients
	message, err := json.Marshal(events.NewBroadcastMessageEvent(server.ID, event.ID, event.Content.Message))
	if err != nil {
		server.Logger.Log(fmt.Sprintf("Error marshalling response: %s\n", err), logger.ERROR)
	} else {
		errs := server.BroadcastMessage(message, conn)
		for _, err := range errs {
			server.Logger.Log(fmt.Sprintf("Error broadcasting message: %s\n", err), logger.ERROR)
		}
	}
}
