package client

import (
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Worker struct {
	Id          string
	Url         string
	rand        *rand.Rand
	BytesMin    int
	BytesMax    int
	IntervalMax int
	IntervalMin int
}

func NewWorker(id string, url string, bytesMin, bytesMax int, intervalMin, intervalMax int) *Worker {
	return &Worker{
		Id:          id,
		Url:         url,
		rand:        rand.New(rand.NewSource(time.Now().UnixNano())),
		BytesMin:    bytesMin,
		BytesMax:    bytesMax,
		IntervalMin: intervalMin,
		IntervalMax: intervalMax,
	}
}

func (w *Worker) getRandomInt(min, max int) int {
	return w.rand.Intn(max-min+1) + min
}

func (w *Worker) sleep(min, max int) {
	i := w.getRandomInt(min, max)
	d := time.Duration(i)
	w.info("Sleep for seconds: " + strconv.Itoa(i))
	fmt.Println("--------------------------------------------------")
	time.Sleep(time.Second * d)
}

func (w *Worker) info(msg string) {
	log.Printf("W%s %s\n", w.Id, msg)
}

func (w *Worker) fatal(msg string) {
	log.Fatalf("W%s %s\n", w.Id, msg)
}

func (w *Worker) Run() {
	for {
		buf := make([]byte, w.getRandomInt(w.BytesMin, w.BytesMax))
		_, err := w.rand.Read(buf)
		if err != nil {
			w.fatal(fmt.Sprintf("error while generating random bytes: %s", err))
			w.sleep(3, 15)
			continue
		}

		w.sendRequest(buf)
		w.sleep(w.IntervalMin, w.IntervalMax)
	}
}

func (w *Worker) sendRequest(bArray []byte) {
	bodyReader := bytes.NewReader(bArray)
	req, err := http.NewRequest(http.MethodPost, w.Url, bodyReader)

	if err != nil {
		w.info(fmt.Sprintf("client: could not create request: %s", err))
		return
	}

	w.info(fmt.Sprintf("Sending %d bytes", len(bArray)))

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		w.info(fmt.Sprintf("client: error making http request: %s", err))
		return
	}
	defer res.Body.Close()

	w.info(fmt.Sprintf("Rsponse:%s", res.Status))
}
