package main

import (
	"github.com/hamednikzad/fakery/internal/server"
	"github.com/hamednikzad/fakery/pkg/shutdown"
)

func main() {
	cfg := getConfig()

	server.Run(cfg.Server.ListenAddress)

	shutdown.Gracefully()
}
