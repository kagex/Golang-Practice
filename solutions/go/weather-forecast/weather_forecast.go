// Package weather provides the weather forecast.
package weather

var (
    // CurrentCondition describes the current weather condition.
    CurrentCondition string
    // CurrentLocation is a location to forecast.
	CurrentLocation  string 
)

//Forecast for current location.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
