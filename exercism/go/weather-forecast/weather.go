/*
Package weather
*/
package weather

// CurrentCondition is ...
var CurrentCondition string

// CurrentLocation is ...
var CurrentLocation string

// Forecast return current weather
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
