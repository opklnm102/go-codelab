# [CodeLab](https://mingrammer.com/go-codelab/codelab-introduction/)

<br>

## Simple Overview
Assume that there are 3 sensors and we have realtime server for handling the data from pipeline

```
|---- Sensor1 ----\            /----- Server1 -----\
                   \          /                     \
|---- Sensor2 ----------------------- Server2 ---------------- Log handler
                   /          \                     /               |
|---- Sensor3 ----/            \----- Server3 -----/                |
         |                               |                          |
[Produce the data]             [Handle the request]         [Logging the data]
```
The sensors send produced data to pipeline concurrently using goroutine, so pipeline should queueing the streaming dta then the server will fetch the data from pipeline by channel and processing with that

<br>

## codelab에서 배울 수 있는 것
- Go application structure
- How to work Go application with packeages
- Go interface/struct model
- Concurrency in Go
- Communication way between goroutines using Go channel

<br>

## 실행 해보기
```sh
$ go run sensor_client.go  # or build

## 다른 shell에서 실행
$ go run sensor_server.go  # or build


## sensor_client
[AtomsphericSensor] Sent : TemperatureSensor : 42.810942, 0.239038
[AtomsphericSensor] Sent : TemperatureSensor : 87.050709, 24.772865
[VelocitySensor] Sent : AccelerometerSensor : 612.156659, 768.804356, 732.840247
...

## sensor_server
[AtomsphericSensor] Received : TemperatureSensor : 42.810942, 0.239038
[VelocitySensor] Received : AccelerometerSensor : 772.477171, 458.903219, 421.887622
[VelocitySensor] Received : GyroSensor : 38.873719 92.231791 178.532700
...
```
