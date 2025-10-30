package data


import (
	"math/rand"
	"time"
)

func GenerateRandomValue(min, max int) float64 {
	return float64(rand.Intn(max-min+1) + min)
}

func GetDaysInYear(year int) []time.Time {
	var dates []time.Time
	start := time.Date(year, 1, 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(1, 0, 0)

	for d := start; d.Before(end); d = d.AddDate(0, 0, 1) {
		dates = append(dates, d)
	}
	return dates
}
