package meteorology

import(
	"fmt"
)

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (t TemperatureUnit) String() string{
	u:= []string{"°C","°F"}
	return u[t]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (t Temperature) String() string{
	return fmt.Sprintf("%d %s",t.degree, t.unit.String())
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (s SpeedUnit) String() string{
	u:= []string{"km/h","mph"}
	return u[s]
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (s Speed) String() string{
	return fmt.Sprintf("%d %s",s.magnitude, s.unit.String())
}

type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (m MeteorologyData) String() string{
	return fmt.Sprintf("%s: %s, Wind %s at %s, %d%s Humidity", m.location, m.temperature.String(), m.windDirection, m.windSpeed.String(), m.humidity, "%")
}
