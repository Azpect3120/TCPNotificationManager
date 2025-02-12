package server

import "net"

type Server struct {
}

func New() *Server {
	net.Listen("tcp", ":8080")
	return &Server{}
}
