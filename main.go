package main

import (
	"fmt"
	"os"
	"strconv"
)

func celsiusToFahrenheit(celsius float64) float64 {
	return (celsius * 9 / 5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
	return (fahrenheit - 32) * 5 / 9
}

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run main.go <CtoF|FtoC> <temperature>")
		os.Exit(1)
	}

	conversionType := os.Args[1]
	temperature, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Invalid input. Please enter a numeric value for temperature.")
		os.Exit(1)
	}

	switch conversionType {
	case "CtoF":
		fmt.Printf("%.2f°C is %.2f°F\n", temperature, celsiusToFahrenheit(temperature))
	case "FtoC":
		fmt.Printf("%.2f°F is %.2f°C\n", temperature, fahrenheitToCelsius(temperature))
	default:
		fmt.Println("Invalid conversion type. Use 'CtoF' for Celsius to Fahrenheit or 'FtoC' for Fahrenheit to Celsius.")
		os.Exit(1)
	}
}
