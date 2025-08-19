package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

func main() {
	// This is the entry point for of the application
	var menuChoice int

	displayMenu()
	fmt.Print("Choice: ")

	if _, err := fmt.Scanln(&menuChoice); err != nil {
		log.Printf("Error:", err)
		return
	}
	clearTerminal()

	switch menuChoice {
	case 1:
		f, err := convertCelToFah()
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			return
		}

		fmt.Printf("Temperature in Fahrenheit: %.2f°F\n", f)
		break
	case 2:
		c, err := convertFahToCel()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("Temperature in Celcius: %.2f°C\n", c)
		break
	}

	fmt.Println("Thank you for using Temperature Converter, Goodbye! 😃")
}

// Takes temperature in °C and converts it to °F
func convertCelToFah() (float32, error) {
	var temperature float32
	fmt.Print("Enter the temperature in Celsius: ")
	if _, err := fmt.Scanln(&temperature); err != nil {
		return 0, fmt.Errorf("Invalid temperature: %v", err)
	}

	fahrenheit := (temperature * 9 / 5) + 32

	return fahrenheit, nil
}

// Takes temperature in °F and converts in to °C
func convertFahToCel() (float32, error) {
	var temperature float32
	fmt.Print("Enter the temperature in Fahrenheit: ")
	if _, err := fmt.Scanln(&temperature); err != nil {
		return 0, fmt.Errorf("Invalid temperature: %v", err)
	}

	celsius := (temperature - 32) * 5 / 9
	return celsius, nil
}

func convertCelToKel() {}

func convertKevToCel() {}

func convertFahToKel() {}

func convertKelToFah() {}

func displayMenu() {
	// Display the menu
	fmt.Println(`
	=== Temperature Converter ===
	1. Celsius to Fahrenheit
	2. Fahrenheit to Celsius
	3. Celsius to Kelvin
	4. Kelvin to Celsius
	5. Fahrenheit to Kelvin
	6. Kelvin to Fahrenheit
	`)
}

func clearTerminal() {
	// clears the terminal

	var cmd *exec.Cmd

	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}

	cmd.Stdout = os.Stdout
	cmd.Run()
}
