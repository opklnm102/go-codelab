package models

import (
	"time"
)

// Sensor is common interface for any sensors
type Sensor interface {
	SendingOutputString() string
	ReceivingOutputString() string
	GenerateSensorData(epsilon float64) Sensor
}

// SensorInfo has common fields for any sensors
type SensorInfo struct {
	Name    string    `json:"name"`
	Type    string    `json:"type"`
	GenTime time.Time `json:"gen_time"`
}
