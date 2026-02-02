package tests

import (
	"gymshark-case-study/internal/service"
	"testing"
)

func TestCalculatePacks(t *testing.T) {
	calculator := service.NewPackCalculator([]int{250, 500, 1000, 2000, 5000})

	result := calculator.CalculatePacks(5000)

	if result.Packs[500] != 1 {
		t.Errorf("Expected 1 pack of 500, got %d", result.Packs)
	}
}
