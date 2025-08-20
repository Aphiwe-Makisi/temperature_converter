package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
)

var temp string

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

		fmt.Printf("%v = %.2f°F\n", temp, f)
		break
	case 2:
		c, err := convertFahToCel()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("%v = %.2f°C\n", temp, c)
		break
	case 3:
		k, err := convertCelToKel()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("%v = %.2fK\n", temp, k)
		break
	case 4:
		c, err := convertKevToCel()
		if err != nil {
			fmt.Print("Error: %v", err)
			return
		}

		fmt.Printf("%v = %.2f°C\n", temp, c)
		break
	case 5:
		k, err := convertFahToKel()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("%v = %.2fK\n", temp, k)
		break
	case 6:
		f, err := convertKelToFah()
		if err != nil {
			fmt.Printf("Error: %v", err)
			return
		}

		fmt.Printf("%v = %.2f°F\n", temp, f)
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
	temp = fmt.Sprintf("%v°C", temperature)

	return fahrenheit, nil
}

// Takes temperature in °F and converts it to °C
func convertFahToCel() (float32, error) {
	var temperature float32
	fmt.Print("Enter the temperature in Fahrenheit: ")
	if _, err := fmt.Scanln(&temperature); err != nil {
		return 0, fmt.Errorf("Invalid temperature: %v", err)
	}

	celsius := (temperature - 32) * 5 / 9
	temp = fmt.Sprintf("%v°F", temperature)

	return celsius, nil
}

// Takes temperature in °C and convert it to K
func convertCelToKel() (float32, error) {
	var temperature float32
	const k float32 = 273.15
	fmt.Print("Enter the temperature in Celsius: ")
	if _, err := fmt.Scanln(&temperature); err != nil {
		return 0, fmt.Errorf("Invalid temperature: %v", err)
	}

	kelvin := temperature + k
	temp = fmt.Sprintf("%v°C", temperature)

	return kelvin, nil
}

// Takes temperature in K and convert it to °C
func convertKevToCel() (float32, error) {
	var temperature float32
	const k float32 = 273.15
	fmt.Print("Enter the temperature in Kelvin: ")
	if _, err := fmt.Scanln(&temperature); err != nil {
		return 0, fmt.Errorf("Invalid temperature: %v", err)
	}

	celsius := temperature - k
	temp = fmt.Sprintf("%vK", temperature)

	return celsius, nil
}

// Takes temperature in °F and convert in to K
func convertFahToKel() (float32, error) {
	var temperature float32
	const k float32 = 273.15
	const f float32 = 32
	fmt.Print("Enter the temperature in Fahrenheit: ")
	if _, err := fmt.Scanln(&temperature); err != nil {
		return 0, fmt.Errorf("Invalid temperature: %v", err)
	}

	kelvin := (temperature-f)*5/9 + k
	temp = fmt.Sprintf("%v°F", temperature)

	return kelvin, nil
}

// Takes temperature in K and convert it to °F
func convertKelToFah() (float32, error) {
	var temperature float32
	const k float32 = 273.15
	const f float32 = 32
	fmt.Print("Enter the temperature in Kelvin: ")
	if _, err := fmt.Scanln(&temperature); err != nil {
		return 0, fmt.Errorf("invalid temperature: %v", err)
	}

	fahrenheit := (temperature-273.15)*9/5 + 32
	temp = fmt.Sprintf("%vK", temperature)

	return fahrenheit, nil
}

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
