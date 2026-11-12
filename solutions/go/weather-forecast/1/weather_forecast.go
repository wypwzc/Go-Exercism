// Package weather provides utilities to track, manage, and report 
// current weather conditions and locations for Goblinocus.
package weather

var (
	// CurrentCondition stores the weather status (e.g., "sunny", "rainy") 
	// of the most recently processed forecast. Users can read or check 
	// this variable to know the latest recorded weather condition.
	CurrentCondition string
	
	// CurrentLocation stores the name of the city or area where the 
	// current weather forecast was last updated. Users can read or reference 
	// this variable to know the active location.
	CurrentLocation  string
)

// Forecast returns a formatted string describing the weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}