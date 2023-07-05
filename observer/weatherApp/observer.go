package weatherApp

import "fmt"

type Observer interface {
	Update(temperature, humidity, pressure float64)
}

type Display struct{}

func (d *Display) Update(temperature, humidity, pressure float64) {
	fmt.Printf("Temperature: %.2f°C\n", temperature)
	fmt.Printf("Humidity: %.2f%%\n", humidity)
	fmt.Printf("Pressure: %.2f hPa\n", pressure)
	fmt.Println("--------------------")
}
