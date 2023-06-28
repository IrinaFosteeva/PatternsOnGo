package ducks

import "fmt"

type RedheadDuck struct {
	Duck
}

func (rhd *RedheadDuck) Display() {
	fmt.Println("I'm a Redhead Duck")
}
