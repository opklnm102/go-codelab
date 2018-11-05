package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/opklnm102/codelab/models"
)

type worker struct {
	ticker      *time.Ticker
	sensor      models.Sensor
	sensorError float64
	serverPort  int
}

// 여러 goroutine에서 동시에 사용하므로 동기화 필요
type counter struct {
	mutex sync.Mutex
	num   int
}

func (c *counter) count() {
	c.mutex.Lock()
	c.num++
	c.mutex.Unlock()
}

func (c *counter) value() int {
	return c.num
}

func sensorWorker(done <-chan struct{}, w worker, c *counter) {
	for {
		select {
		case <-done:
			return
		case <-w.ticker.C:
			sensorData := w.sensor.GenerateSensorData(w.sensorError)
			url := getRequetServerURL(w.serverPort)

			fmt.Println(sensorData.SendingOutputString())

			sendJSONSensorData(url, sensorData)

			c.count()
		}
	}
}

func main() {
	numCPUs := runtime.NumCPU()
	runtime.GOMAXPROCS(numCPUs)

	var sendCounter counter
	var wg sync.WaitGroup
	const numWorkers = 3

	done := make(chan struct{})

	wg.Add(numWorkers)

	workerList := [numWorkers]worker{
		worker{
			ticker:      time.NewTicker(500 * time.Millisecond),
			sensor:      models.GyroSensor{},
			sensorError: 4.0,
			serverPort:  8001,
		},
		worker{
			ticker:      time.NewTicker(500 * time.Millisecond),
			sensor:      models.AccelSensor{},
			sensorError: 12.0,
			serverPort:  8002,
		},
		worker{
			ticker:      time.NewTicker(500 * time.Millisecond),
			sensor:      models.TempSensor{},
			sensorError: 1.5,
			serverPort:  8003,
		},
	}

	for _, w := range workerList {
		go func(w worker) {
			sensorWorker(done, w, &sendCounter)
			wg.Done()
		}(w)
	}

	go func() {
		for {
			if sendCounter.value() > 100 {
				close(done)
				return
			}
		}
	}()

	wg.Wait()

	fmt.Printf("\n[Count: %d] Sending is stopped\n", sendCounter.value())
}

func getRequetServerURL(port int) string {
	urlComponents := []string{"http://127.0.0.1", strconv.Itoa(port)}

	return strings.Join(urlComponents, ":")
}

func sendJSONSensorData(url string, sensorValues models.Sensor) {
	jsonBytes, err := json.Marshal(sensorValues)
	if err != nil {
		log.Fatal("Error occurs when marshaling the sensor values")
	}

	buff := bytes.NewBuffer(jsonBytes)

	resp, err := http.Post(url, "application/json", buff)
	if err != nil || resp.StatusCode != 200 {
		log.Fatal("Error occurs when request the post data")
	}
}
