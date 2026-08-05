package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/petrademia/radix/internal/server"
)

func main() {
	addr := flag.String("addr", ":6379", "listen address")
	flag.Parse()
	s := &server.Server{Addr: *addr}
	if err := s.ListenAndServe(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
