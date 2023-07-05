package ducks

import "fmt"

//стратегия

type FlyBehaviorInterface interface {
	Fly()
}

//контекст

type Duck struct {
	FlyBehavior FlyBehaviorInterface
}

func (d *Duck) SetFlyBehavior(fb FlyBehaviorInterface) {
	d.FlyBehavior = fb
}

func (d *Duck) PerformFly() {
	d.FlyBehavior.Fly()
}

//конкретные стратегии

type FlyWithWings struct{}

func (fw *FlyWithWings) Fly() {
	fmt.Println("I'm flying with wings!")
}

type FlyNoWay struct{}

func (fnw *FlyNoWay) Fly() {
	fmt.Println("I can't fly.")
}
