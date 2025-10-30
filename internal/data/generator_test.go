package data

import (
	"testing"
)

func TestGenerateRandomValue(t *testing.T) {
	min := 10
	max := 20
	val := GenerateRandomValue(min, max)
	if val < float64(min) || val > float64(max) {
		t.Errorf("GenerateRandomValue(%d, %d) = %f; want value between %d and %d", min, max, val, min, max)
	}
}
