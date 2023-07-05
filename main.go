package main

import (
	"Patterns/observer/weatherApp"
	"Patterns/strategy/ducks"
	"fmt"
)

func main() {
	architecture := "observer"

	switch architecture {
	case "strategy":
		ducks.Run()
	case "observer":
		weatherApp.Run()
	default:
		fmt.Println("Неизвестная архитектура")
	}
}
