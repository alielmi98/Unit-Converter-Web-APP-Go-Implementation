package services

// ConvertLength converts length units
func ConvertLength(value float64, fromUnit, toUnit string) float64 {
	unitsToMeters := map[string]float64{
		"mm":   0.001,
		"cm":   0.01,
		"m":    1,
		"km":   1000,
		"inch": 0.0254,
		"foot": 0.3048,
		"yard": 0.9144,
		"mile": 1609.34,
	}
	return (value * unitsToMeters[fromUnit]) / unitsToMeters[toUnit]
}

// ConvertWeight converts weight units
func ConvertWeight(value float64, fromUnit, toUnit string) float64 {
	unitsToGrams := map[string]float64{
		"mg": 0.001,
		"g":  1,
		"kg": 1000,
		"oz": 28.3495,
		"lb": 453.592,
	}
	return (value * unitsToGrams[fromUnit]) / unitsToGrams[toUnit]
}

// ConvertTemperature converts temperature units
func ConvertTemperature(value float64, fromUnit, toUnit string) float64 {
	if fromUnit == toUnit {
		return value
	}
	if fromUnit == "C" {
		if toUnit == "F" {
			return (value * 9 / 5) + 32
		}
		return value + 273.15
	} else if fromUnit == "F" {
		if toUnit == "C" {
			return (value - 32) * 5 / 9
		}
		return (value-32)*5/9 + 273.15
	} else {
		if toUnit == "C" {
			return value - 273.15
		}
		return (value-273.15)*9/5 + 32
	}
}
