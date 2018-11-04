package faker

import (
	"math/rand"
	"time"
)

func getRandValueInRange(rangeSize int, epsilon float64) float64 {
	rand.Seed(time.Now().UTC().UnixNano())

	maxDataWithError := float64(rangeSize) + 2*epsilon

	dataInRange := rand.Float64()*maxDataWithError - epsilon

	return dataInRange
}

func GenerateAngleVelocity(epsilon float64) float64 {
	return getRandValueInRange(360, epsilon)
}

func GenerateGravityAcceleration(epsilon float64) float64 {
	return getRandValueInRange(1023, epsilon)
}

func GenerateTemperature(epsilon float64) float64 {
	return getRandValueInRange(104, epsilon)
}

func GenerateHumidity(epsilon float64) float64 {
	return getRandValueInRange(100, epsilon)
}
