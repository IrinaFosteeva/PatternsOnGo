package ducks

import "fmt"

type RubberDuck struct {
	Duck
}

func (rd *RubberDuck) Display() {
	fmt.Println("I'm a Rubber Duck")
}
