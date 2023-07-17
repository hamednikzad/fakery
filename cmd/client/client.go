package main

import (
	"github.com/hamednikzad/fakery/internal/client"
	"github.com/hamednikzad/fakery/pkg/shutdown"
	"log"
	"strconv"
	"strings"
	"time"
)

func main() {
	cfg := getConfig()

	for _, w := range cfg.Workers {
		bytes := strings.Split(w.Bytes, "-")
		bytesMin, _ := strconv.Atoi(bytes[0])
		bytesMax, _ := strconv.Atoi(bytes[1])

		intervals := strings.Split(w.IntervalSecond, "-")
		intervalMin, _ := strconv.Atoi(intervals[0])
		intervalMax, _ := strconv.Atoi(intervals[1])

		log.Printf("Bytes: %d-%d, Interval in sconds: %d-%d", bytesMin, bytesMax, intervalMin, intervalMax)
		log.Printf("Target address: %s", w.RemoteAddress)
		w1 := client.NewWorker(w.Id, w.RemoteAddress, bytesMin, bytesMax, intervalMin, intervalMax)
		go w1.Run()
		time.Sleep(time.Second * 3)
	}

	shutdown.Gracefully()
}
