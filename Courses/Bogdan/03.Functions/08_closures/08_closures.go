package main

import "fmt"

func createTemperatureAdjuster() (func(change float64) float64, float64) {
	baseTemperature := 90.0

	adjustTemperature := func(change float64) float64 {
		baseTemperature = baseTemperature + change
		return baseTemperature
	}
	return adjustTemperature, baseTemperature
}

func main() {
	adjustTemp, originalTemperature := createTemperatureAdjuster()
	fmt.Printf("Original temperature is %.1f\n", originalTemperature)

	fmt.Printf("Adjust Temp +1.5: %.1fC\n", adjustTemp(1.5)) // baseTemperature is changed
	fmt.Printf("Adjust Temp -3.0: %.1fC\n", adjustTemp(-3.0))

	fmt.Printf("Original temperature is %.1f\n", originalTemperature) // originalTemperature is not changed
}
