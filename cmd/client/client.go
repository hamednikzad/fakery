package main

import (
	"github.com/hamednikzad/fakery/internal/client"
	"github.com/hamednikzad/fakery/pkg/shutdown"
	"time"
)

func main() {
	cfg := getConfig()

	for _, w := range cfg.Workers {
		w1 := client.NewWorker(w.Id, w.RemoteAddress)
		go w1.Run()
		time.Sleep(time.Second * 3)
	}

	shutdown.Gracefully()
}
