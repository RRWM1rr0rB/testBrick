package main

import (
	"fmt"
)

func calculatePressure(row, col int, pressures map[[2]int]float64) float64 {
	if row == 0 && col == 0 {
		return 0
	}

	if pressure, exists := pressures[[2]int{row, col}]; exists {
		return pressure
	}

	pressure := 0.0

	if col > 0 {
		pressure += (calculatePressure(row-1, col-1, pressures) + 1) / 2
	}

	if col < row {
		pressure += (calculatePressure(row-1, col, pressures) + 1) / 2
	}

	pressures[[2]int{row, col}] = pressure
	fmt.Printf("Calculated pressure at (%d, %d): %.5f kg\n", row, col, pressure)
	return pressure
}

func calculateTrianglePressures(rows int) map[[2]int]float64 {
	pressures := make(map[[2]int]float64)

	for row := 0; row < rows; row++ {
		for col := 0; col <= row; col++ {
			_ = calculatePressure(row, col, pressures)
		}
	}
	return pressures
}

func main() {
	rows := 323
	pressures := calculateTrianglePressures(rows)

	// Test outputs with expected values for verification
	testCases := map[[2]int]float64{
		{0, 0}: 0,
		{1, 0}: 0.5, {1, 1}: 0.5,
		{2, 0}: 0.75, {2, 1}: 1.5, {2, 2}: 0.75,
		{3, 0}: 0.875, {3, 1}: 2.125, {3, 2}: 2.125, {3, 3}: 0.875,
		{322, 156}: 306.48749781747574,
	}

	for pos, expected := range testCases {
		calculated := pressures[[2]int{pos[0], pos[1]}]
		fmt.Printf("Pressure at (%d, %d): Calculated = %.5f, Expected = %.5f\n", pos[0], pos[1], calculated, expected)
		if calculated != expected {
			fmt.Printf("Mismatch at (%d, %d): Calculated = %.5f, Expected = %.5f\n", pos[0], pos[1], calculated, expected)
		}
	}
}
