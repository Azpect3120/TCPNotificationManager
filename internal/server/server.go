package server

import (
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
}

// Provide an address for the server to bind to.
func WithAddr(addr string) ServerOptsFunc {
	return func(opts *ServerOpts) {
		opts.Addr = addr
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
