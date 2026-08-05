package server

import "net"

// Server speaks RESP over TCP. YOU IMPLEMENT.
type Server struct {
	Addr string
}

func (s *Server) ListenAndServe() error {
	return errString("server not implemented")
}

func DialSmoke(addr string) (net.Conn, error) {
	return nil, errString("not implemented")
}

type errString string

func (e errString) Error() string { return string(e) }
