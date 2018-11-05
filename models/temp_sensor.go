package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/opklnm102/codelab/faker"
)

// TempSensor produces temperature and humidity values
type TempSensor struct {
	SensorInfo
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

func (s TempSensor) String() string {
	var result []string

	st := fmt.Sprintf("Measured on %s", s.GenTime)
	result = append(result, st)

	st = fmt.Sprintf("Temperature : %f", s.Temperature)
	result = append(result, st)

	st = fmt.Sprintf("Humidity : %f", s.Humidity)
	result = append(result, st)

	return strings.Join(result, "\n")
}

func (s TempSensor) InlineString() string {
	InlineStr := fmt.Sprintf("[Type: %s, Name: %s, Temp: %f, Humidity: %f]",
		s.Type, s.Name, s.Temperature, s.Humidity)

	return InlineStr
}

func (s TempSensor) SendingOutputString() string {
	output := fmt.Sprintf("[%s] Sent : %s : %f, %f",
		s.Type, s.Name, s.Temperature, s.Humidity)

	return output
}

func (s TempSensor) ReceivingOutputString() string {
	output := fmt.Sprintf("[%s] Received : %s : %f, %f",
		s.Type, s.Name, s.Temperature, s.Humidity)

	return output
}

func (s TempSensor) GenerateSensorData(epsilon float64) Sensor {
	tempSensorData := TempSensor{
		SensorInfo: SensorInfo{
			Name:    "TemperatureSensor",
			Type:    "AtomsphericSensor",
			GenTime: time.Now(),
		},
		Temperature: faker.GenerateTemperature(epsilon),
		Humidity:    faker.GenerateHumidity(epsilon),
	}

	return tempSensorData
}
