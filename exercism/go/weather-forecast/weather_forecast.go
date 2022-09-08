//Package weather comment.
package weather

//CurrentCondition comment.
var CurrentCondition string
//CurrentLocation comment.
var CurrentLocation string

//Forecast comment.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
