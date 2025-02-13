package server

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/Azpect3120/TCPNotificationManager/internal/utils"
)

// Function symbol used to configure the server
type ServerOptsFunc func(*ServerOpts)

// Options used to configure the server
type ServerOpts struct {
	// Address to bind
	Addr string

	// Port for the server to listen on
	Port int

	// Max connection limit, will throw an error if exceeded
	MaxConn int

	// Use TLS to secure the connection
	TLS bool

	// Size of the message buffer in bytes
	MsgBufSize int
}

// Provide an address for the server to bind to.
func WithAddr(addr string) ServerOptsFunc {
	return func(opts *ServerOpts) {
		opts.Addr = addr
	}
}

// Use TLS to secure the connection.
func WithTLS() ServerOptsFunc {
	return func(opts *ServerOpts) {
		opts.TLS = true
	}
}

// Provide an port for the server to listen on.
func WithPort(port int) ServerOptsFunc {
	return func(opts *ServerOpts) {
		opts.Port = port
	}
}

// Provide a max connection limit for the server.
func WithMaxConn(maxConn int) ServerOptsFunc {
	return func(opts *ServerOpts) {
		opts.MaxConn = maxConn
	}
}

// Provide a message buffer size for the server.
func WithMsgBufSize(msgBufSize int) ServerOptsFunc {
	return func(opts *ServerOpts) {
		opts.MsgBufSize = msgBufSize
	}
}

// Defines the default server options, if they are not
// provided by the user.
func defaultServerOpts() ServerOpts {
	return ServerOpts{
		Addr:       "127.0.0.1",
		Port:       8080,
		TLS:        false,
		MaxConn:    10,
		MsgBufSize: 1024,
	}
}

// EventHandler is a function type that is used to handle events
// that are received by the server. This is built using a generic
// to accept any type of event.
//
// This type will define the function signature for the event
// handlers that are used by the server.
type EventHandler[T any] func(net.Conn, *T)

// TcpServer is a struct that represents a TCP server.
// Server options are abstracted away from the user in a
// way that they can be provided as arguments to the NewTCPServer
// function.
//
// For an explanation of this pattern, see the following video:
// https://www.youtube.com/watch?v=MDy7JQN5MN4&t=601s
type TcpServer struct {
	// Server options.
	Opts ServerOpts

	// ID of the server.
	ID string

	// Connections to the server. A slice of connections. This
	// slice should not be used for authentication, but for
	// simply holding the connections (both authenticated and
	// unauthenticated).
	//
	// When the length of this slice becomes equal to the max
	// connections limit defined in the server options, the
	// server will no longer accept connections.
	Conns []net.Conn

	// A list of authorized clients. The client ID is the key,
	// and the client's connection is the value. This is used to
	// verify that the client ID is being used by the correct
	// client. This is a simple way to verify the client's
	// identity.
	//
	// This map will be used to determine which clients to publish
	// messages to. If the client is not in this map, but is in the
	// Conns slice, the server will not publish messages to that
	// connection.
	Authorized map[string]net.Conn

	// Store any errors that occur during the server's lifecycle.
	Errors []error

	// TLS configuration for the server.
	TLSConfig *tls.Config

	// EventHandlers is a map of event types to their handlers. This map
	// will be used to determine which function to call when an event is
	// received by the server.
	//
	// The key is the event name, and the value is the function to call.
	// The function is defined here as an interface{} but it should be
	// defined as the EventHandler type.
	EventHandlers map[string]interface{}
}

// RegisterEventHandler registers an event handler for a specific event type.
// Methods cannot have generic types, so this function will be used to register
// the event handlers for the server.
//
// The event name should match the class of event that is being handled, not the
// name of the event stored in the struct. This is because we use reflection to
// determine which handler to call.
func RegisterEventHandler[T any](server *TcpServer, eventName string, handler EventHandler[T]) {
	server.EventHandlers[eventName] = handler

}

// Create a new TCP server with the provided options. If options
// are not provided, the default options will be used. They can
// be found in the defaultServerOpts function.
func NewTCPServer(opts ...ServerOptsFunc) *TcpServer {
	server := &TcpServer{
		Opts: defaultServerOpts(),
		ID:   utils.GenerateServerID(),
	}

	// Apply the options to the server.
	for _, optFn := range opts {
		optFn(&server.Opts)
	}

	// Create the connections slice here using the max connection limit.
	// This could be done in the instantiating of the server, but it is
	// done here to show that the server is created with a max connection
	// limit.
	server.Conns = make([]net.Conn, 0, server.Opts.MaxConn)
	server.Authorized = make(map[string]net.Conn)

	// Initialize the event handlers map
	server.EventHandlers = make(map[string]interface{})
	RegisterEventHandler(server, "RequestAuthenticationEvent", RequestAuthenticationHandler)

	return server
}

// This function is used to configure the TLS settings for the server.
// It is only required if you intend to use TLS, otherwise, there is no
// need to run this function.
//
// The server is returned to allow for method chaining.
//
// Errors will not be handled here, they will be stored in the server's
// errors slice. If the files do not exist, or cannot be read, an error
// will be stored in the server's errors slice.
func (s *TcpServer) Configure(certPath, keyPath string) *TcpServer {
	// Ensure the files exist
	if _, err := os.OpenFile(certPath, os.O_RDONLY, 0644); err != nil {
		s.Errors = append(s.Errors, err)
	}
	if _, err := os.OpenFile(keyPath, os.O_RDONLY, 0644); err != nil {
		s.Errors = append(s.Errors, err)
	}

	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		s.Errors = append(s.Errors, err)
	}

	// TODO: Switch these values to correct production values
	s.TLSConfig = &tls.Config{
		Certificates: []tls.Certificate{cert},
		// Verifying isn't working right now, need to get certs signed by CA
		ClientAuth:         tls.RequestClientCert, // Use tls.RequireAndVerifyClientCert for production and security
		InsecureSkipVerify: true,                  // ONLY FOR TESTING - NEVER IN PRODUCTION
	}

	return s
}

// Listen starts the server and listens for incoming connections. For a
// server that uses TLS, the server will use the TLS configuration provided.
// Make sure to call the Configure function before calling Listen if you
// intend to use TLS. Otherwise, the server will not use TLS.
//
// This is a result of a possible null pointer deference if the TLS configuration
// is not set before the server starts listening.
//
// The listener object is returned by this function and can be used by the caller.
// The caller is the owner of the memory and is responsible for closing the listener.
func (s *TcpServer) Listen() net.Listener {
	var ln net.Listener
	var err error

	if s.Opts.TLS && s.TLSConfig != nil {
		ln, err = tls.Listen("tcp", fmt.Sprintf("%s:%d", s.Opts.Addr, s.Opts.Port), s.TLSConfig)
	} else {
		ln, err = net.Listen("tcp", fmt.Sprintf("%s:%d", s.Opts.Addr, s.Opts.Port))
	}

	if err != nil {
		s.Errors = append(s.Errors, err)
	}
	return ln
}

// Add a connection to the server. This function does not authenticate the
// client, but it does allow the server to track the connection. If the server
// is at the max connection limit, the server will return an error.
func (s *TcpServer) addConnection(conn net.Conn) error {
	// At max size, return an error. The error will be used to send back a connection_rejected
	// message to the client.
	if len(s.Conns) > s.Opts.MaxConn {
		return fmt.Errorf("Max connection limit reached")
	}

	s.Conns = append(s.Conns, conn)
	return nil
}

// Remove a connection from the server. This function should be called when the
// connection closes.
//
// It will also remove the connection from the authorized map
// once the map is implemented.
func (s *TcpServer) removeConnection(conn net.Conn) {
	for i, c := range s.Conns {
		if c == conn {
			s.Conns = append(s.Conns[:i], s.Conns[i+1:]...)
			break
		}
	}
}
