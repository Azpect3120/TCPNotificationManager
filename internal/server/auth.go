package server

import "net"

// Check if the client is authenticated. If they are not authenticated, they
// will be able to send messages to the server and they will not be able to
// receive messages from the server.
//
// Using the clientID and the address of the client, the server can determine
// if the client is authenticated and if the client ID is being used by
// another client.
//
// This function assumes the address stored in the server's Authorized map
// is the RemoteAddr of the client.
func (s *TcpServer) isAuthenticated(clientID string, conn net.Conn) bool {
	// Get the connection from the authorized map via the clientID.
	// If the client is not in the map, they are not authenticated.
	connAuth, ok := s.Authorized[clientID]
	if !ok {
		return false
	}

	// If they are in the map, validate that the connection sending the message
	// is the same connection that was authorized to use the clientID.
	return conn.RemoteAddr().String() == connAuth.RemoteAddr().String()
}
