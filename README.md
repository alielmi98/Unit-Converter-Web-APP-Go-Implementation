# Unit Converter Project

Welcome to the Unit Converter project! This is a beginner-level web application designed to help users convert between different units of measurement. It's part of the roadmap.sh community projects aimed at assisting learners and developers in honing their skills.

This project is tailored for backend development and is suitable for beginners looking to enhance their programming skills. It is hosted on [roadmap.sh](https://roadmap.sh/projects/unit-converter) and is designed for educational purposes.

## Overview

The Unit Converter allows users to input a value and select the units they wish to convert from and to. The application supports conversions for various types of measurements, including:

- **Length**: millimeter, centimeter, meter, kilometer, inch, foot, yard, mile
- **Weight**: milligram, gram, kilogram, ounce, pound
- **Temperature**: Celsius, Fahrenheit, Kelvin

## Project Structure

```
unit-converter/
├── main.go
├── config/
│ └── config.go
├── handlers/
│ ├── length.go
│ ├── temperature.go
│ ├── weight.go
│ └── index.go
├── services/
│ └── converter.go
├── templates/
│ ├── index.html
│ ├── length.html
│ ├── weight.html
│ └── temperature.html
└── static/
└── styles.css
```

## Features

- User-friendly web interface for inputting values and selecting units.
- Instant conversion results displayed on the same page.
- No database required; the application processes data in real-time.

## Requirements

- Build a simple web page with sections for each type of unit conversion.
- Allow users to input values, select units for conversion, and view the results.
- The application should handle conversions without needing a backend database.

## How It Works

1. **Input Value**: Users enter the value they wish to convert.
2. **Select Units**: Users choose the units to convert from and to.
3. **Display Result**: Upon submission, the application calculates the converted value and displays it on the webpage.

The application consists of three main pages for different unit conversions: length, weight, and temperature. Each page includes a form for inputting values and selecting units.

## Getting Started

To get started with the Unit Converter project, follow these steps:

1. Clone the repository:
```bash
git clone https://github.com/alielmi98/Unit-Converter-Web-APP-Go-Implementation.git
```
2. Navigate to the project directory:
```bash
cd unit-converter
```
3. Open the `index.html` file in your web browser to view the application.

## Contribution

This project is designed for educational purposes and is open for contributions. If you have suggestions for improvements or new features, feel free to submit a pull request!

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
