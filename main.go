package main

import (
	"Patterns/strategy/ducks"
	"fmt"
)

func main() {
	architecture := "strategy"

	switch architecture {
	case "strategy":
		ducks.Run()
	default:
		fmt.Println("Неизвестная архитектура")
	}
}
