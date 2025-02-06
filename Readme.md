# Temperature Converter in Go

This is a simple command-line application written in Go that converts temperatures between Celsius and Fahrenheit.

## Features

- Converts Celsius to Fahrenheit
- Converts Fahrenheit to Celsius
- Provides a simple command-line interface

## Prerequisites

- Go installed on your system ([Download Go](https://go.dev/dl/))

## Installation

Clone this repository and navigate into the project directory:

```sh
git clone https://github.com/jkimunyi-dev/temperature_converter_go.git
cd temperature_converter_go
```

## Usage

Run the program using the following command:

```sh
go run main.go <CtoF|FtoC> <temperature>
```

### Examples:

Convert **25°C to Fahrenheit**:

```sh
go run main.go CtoF 25
```

**Output:**

```
25.00°C is 77.00°F
```

Convert **77°F to Celsius**:

```sh
go run main.go FtoC 77
```

**Output:**

```
77.00°F is 25.00°C
```

## Error Handling

- If incorrect arguments are provided, the program will display a usage message.
- If the temperature value is not a valid number, an error message is displayed.

## License

This project is open-source and available under the MIT License.
