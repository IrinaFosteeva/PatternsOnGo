package ducks

import "fmt"

type Duck struct {
	FlyBehavior FlyBehavior
}

func (d *Duck) SetFlyBehavior(fb FlyBehavior) {
	d.FlyBehavior = fb
}

func (d *Duck) PerformFly() {
	d.FlyBehavior.Fly()
}

type FlyBehavior interface {
	Fly()
}

type FlyWithWings struct{}

func (fw *FlyWithWings) Fly() {
	fmt.Println("I'm flying with wings!")
}

type FlyNoWay struct{}

func (fnw *FlyNoWay) Fly() {
	fmt.Println("I can't fly.")
}
