package cars

const carsPerHour int = 221

// CalculateProductionRatePerHour for the assembly line, taking into account
// its success rate
func CalculateProductionRatePerHour(speed int) float64 {
	return successRate(speed) * float64(speed) * float64(carsPerHour)
}

// CalculateProductionRatePerMinute describes how many working items are
// produced by the assembly line every minute
func CalculateProductionRatePerMinute(speed int) int {
	return int(CalculateProductionRatePerHour(speed) / 60)
}

// successRate is used to calculate the ratio of an item being created without
// error for a given speed
func successRate(speed int) float64 {
	var rate float64

	switch {
	case speed > 0 && speed < 5:
		rate = float64(1)
	case speed < 9:
		rate = float64(.9)
	case speed < 11:
		rate = float64(.77)
	}

	return rate
}
