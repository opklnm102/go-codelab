package models

import (
	"fmt"
	"strings"
	"time"

	"github.com/opklnm102/codelab/faker"
)

func (s GyroSensor) String() string {
	var result []string

	st := fmt.Sprintf("Measured on %s", s.GenTime)
	result = append(result, st)

	st = fmt.Sprintf("Angle Velocity of X-axis : %f", s.AngleVelocityX)
	result = append(result, st)

	st = fmt.Sprintf("Angle Velocity of Y-axis : %f", s.AngleVelocityY)
	result = append(result, st)

	st = fmt.Sprintf("Angle Velocity of Z-axis : %f", s.AngleVelocityZ)
	result = append(result, st)

	return strings.Join(result, "\n")
}

// GyroSensor produces x-y-z axes angle velocity values
type GyroSensor struct {
	SensorInfo
	AngleVelocityX float64 `json:"x_axis_angle_velocity"`
	AngleVelocityY float64 `json:"y_axis_angle_velocity"`
	AngleVelocityZ float64 `json:"z_axis_angle_velocity"`
}

// SendingOutputString of GyroSensor create sending output string with their data
func (s GyroSensor) SendingOutputString() string {
	output := fmt.Sprintf("[%s] Sent : %s : %f, %f, %f",
		s.Type, s.Name, s.AngleVelocityX, s.AngleVelocityY, s.AngleVelocityZ)

	return output
}

func (s GyroSensor) ReceivingOutputString() string {
	output := fmt.Sprintf("[%s] Received : %s : %f %f %f",
		s.Type, s.Name, s.AngleVelocityX, s.AngleVelocityY, s.AngleVelocityZ)

	return output
}

func (s GyroSensor) GenerateSensorData(epsilon float64) Sensor {
	GyroSensorData := GyroSensor{
		SensorInfo: SensorInfo{
			Name:    "GyroSensor",
			Type:    "VelocitySensor",
			GenTime: time.Now(),
		},
		AngleVelocityX: faker.GenerateAngleVelocity(epsilon),
		AngleVelocityY: faker.GenerateAngleVelocity(epsilon),
		AngleVelocityZ: faker.GenerateAngleVelocity(epsilon),
	}

	return GyroSensorData
}
