package server

import (
	"errors"
	"fmt"
	"io"
	"net"
	"reflect"

	"github.com/Azpect3120/TCPNotificationManager/internal/events"
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

		// This is where the messages should be parsed and processed.
		if n > 0 {
			event, err := events.Parser(buf[:n])
			if err != nil {
				// Not sure why or when this would happen
				fmt.Printf("Error parsing message: %v\n", err)
			}

			// Handle the event. A check for authorization should be done
			// in the handlers for the events, because there is no way to
			// get data from the raw interface{} type until it has been
			// type asserted.
			eventType := reflect.TypeOf(event).Elem()
			eventName := eventType.Name()

			if handler, ok := s.EventHandlers[eventName]; ok {
				// Correctly assert the handler type
				handlerType := reflect.TypeOf(handler)

				// Check if the handler type matches the event type. This
				// shit is black magic, Gemini created it for me, it seems
				// make sense but definitely not something I could have done
				// on my own.
				if handlerType.NumIn() == 3 && handlerType.In(2) == reflect.PointerTo(reflect.TypeOf(event).Elem()) {
					reflect.ValueOf(handler).Call([]reflect.Value{reflect.ValueOf(s), reflect.ValueOf(conn), reflect.ValueOf(event)})
				} else {
					fmt.Println("Handler type mismatch for", eventName)
				}
			} else {
				fmt.Println("No handler found for", eventName)
			}
		}
	}
}
