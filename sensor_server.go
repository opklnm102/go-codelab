package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/opklnm102/codelab/models"
)

const (
	logDir   = "log"
	tempLog  = "Temp.log"
	accelLog = "Accel.log"
	gyroLog  = "Gyro.log"
)

type LogContent struct {
	content    string
	location   string
	sensorName string
}

type GyroHandler struct {
	buf chan<- LogContent
}

type AccelHandler struct {
	buf chan<- LogContent
}

type TempHandler struct {
	buf chan<- LogContent
}

func (m *TempHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var data models.TempSensor

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error Occurred when parsing temperature data")
	}
	defer req.Body.Close()

	fmt.Println(data.ReceivingOutputString())

	m.buf <- LogContent{
		content:    fmt.Sprintf("%s", data),
		location:   tempLog,
		sensorName: data.Name,
	}
}

func (m *GyroHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var data models.GyroSensor

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error Occurred when parsing gyroscope data")
	}
	defer req.Body.Close()

	fmt.Println(data.ReceivingOutputString())

	m.buf <- LogContent{
		content:    fmt.Sprintf("%s", data),
		location:   gyroLog,
		sensorName: data.Name,
	}
}

func (m *AccelHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	var data models.AccelSensor

	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&data)
	if err != nil {
		fmt.Println("Error Occurred when parsing accelerator data")
	}
	defer req.Body.Close()

	fmt.Println(data.ReceivingOutputString())

	m.buf <- LogContent{
		content:    fmt.Sprintf("%s", data),
		location:   accelLog,
		sensorName: data.Name,
	}
}

func fileLogger(logStream <-chan LogContent) {
	dir, _ := os.Open("log")
	dirInfo, _ := dir.Stat()

	if dirInfo == nil {
		err := os.Mkdir("log", os.ModePerm)

		if err != nil {
			log.Fatal("Error creating directory 'log'\n", err)
		}
	}
	dir.Close()

	for logData := range logStream {
		joinee := []string{logDir, logData.location}
		filePath := strings.Join(joinee, "/")

		fileHandle, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

		if err != nil {
			log.Fatal("Error opening file\n", err)
		}

		logger := log.New(fileHandle, "", log.LstdFlags)
		logger.Printf("[%s Data Received]\n%s\n", logData.sensorName, logData.content)

		defer fileHandle.Close()
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(4)

	logBuf := make(chan LogContent)
	gyroHandler := &GyroHandler{buf: logBuf}
	accelHandler := &AccelHandler{buf: logBuf}
	tempHandler := &TempHandler{buf: logBuf}

	go http.ListenAndServe(":8001", gyroHandler)
	go http.ListenAndServe(":8002", accelHandler)
	go http.ListenAndServe(":8003", tempHandler)
	go fileLogger(logBuf)

	wg.Wait()
}
