# 🌡️ Temperature Converter

A simple command-line **temperature converter** written in **Go** that converts between **Celsius, Fahrenheit, and Kelvin** temperature scales.  

---

## ✨ Features

- Multiple conversions supported:
  - Celsius ↔ Fahrenheit
  - Celsius ↔ Kelvin
  - Fahrenheit ↔ Kelvin
- Interactive menu interface
- Input validation and error handling
- Cross-platform terminal clearing
- Precise calculations using `float32`

---

## ⚙️ Installation

### Prerequisites
- Go **1.16+**
- Terminal / Command Prompt access

### Setup
```bash
# Clone or download the repository
git clone <repo-url>
cd temperature-converter

# Build the application
go build -o temperature-converter main.go
🚀 Usage
Run the application
./temperature-converter
Or run directly with Go:

go run main.go
Menu Options
When you run the application, you’ll see:

=== Temperature Converter ===

1. Celsius to Fahrenheit
2. Fahrenheit to Celsius
3. Celsius to Kelvin
4. Kelvin to Celsius
5. Fahrenheit to Kelvin
6. Kelvin to Fahrenheit

Example
Choice: 1
Enter the temperature in Celsius: 25
25°C = 77.00°F

Thank you for using Temperature Converter, Goodbye! 😃
🔢 Conversion Formulas
Celsius → Fahrenheit: °F = (°C × 9/5) + 32

Fahrenheit → Celsius: °C = (°F - 32) × 5/9

Celsius → Kelvin: K = °C + 273.15

Kelvin → Celsius: °C = K - 273.15

Fahrenheit → Kelvin: K = (°F - 32) × 5/9 + 273.15

Kelvin → Fahrenheit: °F = (K - 273.15) × 9/5 + 32

📂 Code Structure
main.go

main() – Entry point & menu handling

displayMenu() – Shows conversion options

clearTerminal() – Clears terminal (cross-platform)

convertCelToFah() – Celsius → Fahrenheit

convertFahToCel() – Fahrenheit → Celsius

convertCelToKel() – Celsius → Kelvin

convertKevToCel() – Kelvin → Celsius

convertFahToKel() – Fahrenheit → Kelvin

convertKelToFah() – Kelvin → Fahrenheit

🛡️ Error Handling
Invalid menu selections

Non-numeric temperature inputs

Input/output errors during user interaction

💻 Platform Compatibility
Windows → uses cmd /c cls

Unix/Linux/macOS → uses clear

🔧 Technical Details
Language: Go (Golang)

Input type: float32

Precision: 2 decimal places

Dependencies: Standard library only