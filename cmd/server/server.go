package main

import (
	"github.com/hamednikzad/fakery/internal/server"
	"github.com/hamednikzad/fakery/pkg/shutdown"
)

const colorRed = "\033[0;31m"
const colorNone = "\033[0m"

func main() {
	server.Run(":5000")

	shutdown.Gracefully()
}
