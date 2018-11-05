package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/opklnm102/codelab/faker"
)

// AccelSensor produces x-y-z axes gravity acceleration values
type AccelSensor struct {
	SensorInfo
	GravityAccX float64 `json:"x_axis_gravity_acceleration"`
	GravityAccY float64 `json:"y_axis_gravity_acceleration"`
	GravityAccZ float64 `json:"z_axis_gravity_acceleration"`
}

func (s AccelSensor) String() string {
	var result []string

	st := fmt.Sprintf("Measured on %s", s.GenTime)
	result = append(result, st)

	st = fmt.Sprintf("Gravitational Velocity of X-axis : %f", s.GravityAccX)
	result = append(result, st)

	st = fmt.Sprintf("Gravitational Velocity of Y-axis : %f", s.GravityAccY)
	result = append(result, st)

	st = fmt.Sprintf("Gravitational Velocity of Z-axis : %f", s.GravityAccZ)
	result = append(result, st)

	return strings.Join(result, "\n")
}

func (s AccelSensor) SendingOutputString() string {
	output := fmt.Sprintf("[%s] Sent : %s : %f, %f, %f",
		s.Type, s.Name, s.GravityAccX, s.GravityAccY, s.GravityAccZ)

	return output
}

func (s AccelSensor) ReceivingOutputString() string {
	output := fmt.Sprintf("[%s] Received : %s : %f, %f, %f",
		s.Type, s.Name, s.GravityAccX, s.GravityAccY, s.GravityAccZ)

	return output
}

func (s AccelSensor) GenerateSensorData(epsilon float64) Sensor {
	accelSensorData := AccelSensor{
		SensorInfo: SensorInfo{
			Name:    "AccelerometerSensor",
			Type:    "VelocitySensor",
			GenTime: time.Now(),
		},
		GravityAccX: faker.GenerateGravityAcceleration(epsilon),
		GravityAccY: faker.GenerateGravityAcceleration(epsilon),
		GravityAccZ: faker.GenerateGravityAcceleration(epsilon),
	}

	return accelSensorData
}
