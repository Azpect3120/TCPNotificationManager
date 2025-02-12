package server

import (
	"crypto/tls"
	"fmt"
	"net"
	"os"

	"github.com/Azpect3120/TCPNotificationManager/internal/client"
	"github.com/Azpect3120/TCPNotificationManager/internal/utils"
)

type ServerOptsFunc func(*ServerOpts)

type ServerOpts struct {
	// Address to bind
	Addr string

	// Port for the server to listen on
	Port int

	// Max connection limit, will throw an error if exceeded
	MaxConn int

	// Use TLS to secure connection
	TLS bool
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

// Defines the default server options, if they are not
// provided by the user.
func defaultServerOpts() ServerOpts {
	return ServerOpts{
		Addr:    "127.0.0.1",
		Port:    8080,
		MaxConn: 10,
	}
}

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

	// Connections to the server. The key is the client ID,
	// and the value is the client itself.
	Conns map[string]client.TCPClient

	// Store any errors that occur during the server's lifecycle.
	Errors []error

	// TLS configuration for the server.
	TLSConfig *tls.Config
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

	// Create the connections map here using the max connection limit.
	// This could be done in the instantiation of the server, but it is
	// done here to show that the server is created with a max connection
	// limit.
	server.Conns = make(map[string]client.TCPClient, server.Opts.MaxConn)

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
		Certificates:       []tls.Certificate{cert},
		ClientAuth:         tls.NoClientCert, // Use tls.RequireAndVerifyClientCert for production and security
		InsecureSkipVerify: true,             // Use false for production and safety
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
