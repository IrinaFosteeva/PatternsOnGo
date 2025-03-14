package weatherApp

import "fmt"///test
func Run() {
	fmt.Println("Running observer architecture...")

	weatherData := &WeatherData{}

	display := &Display{}
	weatherData.RegisterObserver(display)

	// Код для получения и установки значений температуры, влажности и давления

	weatherData.SetMeasurements(25.0, 60.0, 1013.25)
}
