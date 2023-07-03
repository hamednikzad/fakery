package main

import (
	"github.com/hamednikzad/fakery/internal/server"
	"github.com/hamednikzad/fakery/pkg/shutdown"
)

func main() {
	server.Run(":5000")

	shutdown.Gracefully()
}
