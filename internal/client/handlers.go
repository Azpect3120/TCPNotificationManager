package client

import (
	"fmt"

	"github.com/Azpect3120/TCPNotificationManager/internal/events"
	"github.com/Azpect3120/TCPNotificationManager/internal/logger"
)

// Handle the ConnectionAcceptedEvent sent by the server to the client. This
// event is sent when the server accepts the connection from the client. All
// this function must do is update the client with the ID generated by the
// server and returned in the event.
func ConnectionAcceptedHandler(client *TcpClient, event *events.ConnectionAcceptedEvent) {
	client.ID = event.Content.ClientID
	client.Logger.Log(fmt.Sprintf("Client ID set to: %s\n", client.ID), logger.DEBUG)
}

// Handle the ClientAuthenticatedEvent sent by the server to the client. This
// event is sent when the server has authenticated the client. This function
// does not really do anything important, but it prints debug messages.
//
// TODO: Implement UI features here.
func ClientAuthenticatedHandler(client *TcpClient, event *events.ClientAuthenticatedEvent) {
	msg := fmt.Sprintf("New client authenticated: %s\n", event.Content.ClientID)
	client.Logger.Log(msg, logger.INFO)
}

// Handle the ClientDisconnectedEvent sent by the server to the client. This
// event is sent when the server has authenticated the client. This function
// does not really do anything important, but it prints debug messages.
//
// TODO: Implement UI features here.
func ClientDisconnectedHandler(client *TcpClient, event *events.ClientDisconnectedEvent) {
	msg := fmt.Sprintf("Client disconnected: %s\n", event.Content.ClientID)
	client.Logger.Log(msg, logger.INFO)
}

// Handle the BroadcastMessageEvent sent by the server to the client. This
// event is sent when the server has authenticated the client. This function
// does not really do anything important, but it prints debug messages.
//
// TODO: Implement UI features here.
func BroadcastMessageHandler(client *TcpClient, event *events.BroadcastMessageEvent) {
	msg := fmt.Sprintf("(%s): %s\n", event.Content.Sender, event.Content.Message)
	client.Logger.Log(msg, logger.INFO)
}
