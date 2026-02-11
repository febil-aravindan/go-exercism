//Package weather forcasts weather of the city.
package weather

var (
//CurrentCondition represents the current condition and CurrentLocation is the current location.
	CurrentCondition string
//CurrentLocation represents the current location. 
	CurrentLocation  string
)

//Forecast funtion will return the given cities weather.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
