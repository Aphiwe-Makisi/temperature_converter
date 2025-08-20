Temperature Converter 🌡️
A simple command-line temperature converter written in Go that converts between Celsius, Fahrenheit, and Kelvin temperature scales.
Features

Multiple conversions supported:

Celsius ↔ Fahrenheit
Celsius ↔ Kelvin
Fahrenheit ↔ Kelvin


Interactive menu interface
Input validation and error handling
Cross-platform terminal clearing
Precise calculations with float32

Installation
Prerequisites

Go 1.16 or higher installed on your system
Terminal/Command prompt access

Setup

Clone or download the repository
Navigate to the project directory
Build the application:
bashgo build -o temperature-converter main.go


Usage
Running the Application
bash# If you built the binary
./temperature-converter

# Or run directly with Go
go run main.go
Menu Options
When you run the application, you'll see an interactive menu:
=== Temperature Converter ===
1. Celsius to Fahrenheit
2. Fahrenheit to Celsius
3. Celsius to Kelvin
4. Kelvin to Celsius
5. Fahrenheit to Kelvin
6. Kelvin to Fahrenheit
Example Usage
Choice: 1
Enter the temperature in Celsius: 25
25°C = 77.00°F
Thank you for using Temperature Converter, Goodbye! 😃
Conversion Formulas
The application uses the following standard temperature conversion formulas:
FromToFormulaCelsiusFahrenheit°F = (°C × 9/5) + 32FahrenheitCelsius°C = (°F - 32) × 5/9CelsiusKelvinK = °C + 273.15KelvinCelsius°C = K - 273.15FahrenheitKelvinK = (°F - 32) × 5/9 + 273.15KelvinFahrenheit°F = (K - 273.15) × 9/5 + 32
Code Structure
main.go
├── main()                 # Entry point and menu handling
├── displayMenu()          # Shows conversion options
├── clearTerminal()        # Cross-platform terminal clearing
├── convertCelToFah()      # Celsius to Fahrenheit conversion
├── convertFahToCel()      # Fahrenheit to Celsius conversion
├── convertCelToKel()      # Celsius to Kelvin conversion
├── convertKevToCel()      # Kelvin to Celsius conversion
├── convertFahToKel()      # Fahrenheit to Kelvin conversion
└── convertKelToFah()      # Kelvin to Fahrenheit conversion
Error Handling
The application includes robust error handling for:

Invalid menu selections
Non-numeric temperature inputs
Input/output errors during user interaction

Platform Compatibility

Windows: Uses cmd /c cls for terminal clearing
Unix/Linux/macOS: Uses clear command for terminal clearing

Technical Details

Language: Go (Golang)
Input Type: float32 for temperature values
Precision: Results displayed to 2 decimal places
Dependencies: Only Go standard library packages

Future Enhancements
Potential improvements for future versions:

Support for additional temperature scales (Rankine, Réaumur)
Batch conversion mode
Configuration file for default settings
Unit tests
More detailed input validation (e.g., absolute zero limits)