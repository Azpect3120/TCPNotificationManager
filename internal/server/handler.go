package server

import (
	"errors"
	"fmt"
	"io"
	"net"
)

// Handle a connection from a client. This method is defined on the
// TcpServer struct, so it can reference the server's state.
//
// The clients must be authenticated before they can be used. But the
// connections do not need to be authenticated, but the messages sent
// by clients should contain the client ID which can be used to
// authenticate and verify the client.
//
// This function will handle the memory management of the connection,
// and will close the connection when it is done.
func (s *TcpServer) HandleConnection(conn net.Conn) {
	// Defer the closing of the connection until the function returns.
	defer func() {
		conn.Close()
		fmt.Printf("Connection lost: %s\n", conn.RemoteAddr().String())
		s.removeConnection(conn)
	}()

	// Add the connection to the server's connection slice. This action
	// does not authenticate the client, but it does allow the server to
	// track the connection.
	if err := s.addConnection(conn); err != nil {
		fmt.Printf("Error adding connection: %s\n", err)
		return
	}

	fmt.Printf("Connection accepted: %s\n", conn.RemoteAddr().String())

	// Create a buffer to read the messages from the clients. The size
	// of the buffer is defined in the server's options. Default is 1KB.
	buf := make([]byte, s.Opts.MsgBufSize)
	for {
		n, err := conn.Read(buf)
		// Connection was closed by the client
		if errors.Is(err, net.ErrClosed) || errors.Is(err, io.EOF) {
			return
		} else if err != nil {
			// Else, a real error occurred
			fmt.Printf("Error reading from connection: %v\n", err)
			return
		}

		if n > 0 {
			// This is where the messages should be parsed and processed.
			fmt.Printf("Read: %s", buf[:n])
		}
	}
}
