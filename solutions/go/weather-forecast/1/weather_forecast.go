// Package weather provide functionality of forecasting the current weather condition of a location.
package weather

var (
	// CurrentCondition represents the current weather condition.
	CurrentCondition string
	// CurrentLocation represents the current location of th weather forecast.
	CurrentLocation string
)

// Forecast takes in city and condition string parameters and returns a string describing the current weather condition of the current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
