package ducks

import "fmt"

func Run() {
	fmt.Println("Running strategy architecture...")

	rubberDuck := &RubberDuck{}
	rubberDuck.SetFlyBehavior(&FlyNoWay{})

	rubberDuck.Display()
	rubberDuck.PerformFly()

	redheadDuck := &RedheadDuck{}
	redheadDuck.SetFlyBehavior(&FlyWithWings{})

	redheadDuck.Display()
	redheadDuck.PerformFly()
}
