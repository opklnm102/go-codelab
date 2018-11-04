# [CodeLab](https://mingrammer.com/go-codelab/codelab-introduction/)


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


## codelab에서 배울 수 있는 것
- Go application structure
- How to work Go application with packeages
- Go interface/struct model
- Concurrency in Go
- Communication way between goroutines using Go channel
