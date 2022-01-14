package main

import "fmt"

type sports interface {
	turnOnTurbo()
}

type bwm struct {
	engine     string
	isTurbosOn bool
}

func (b *bwm) turnOnTurbo() {
	b.isTurbosOn = true
}

func main() {
	car := bwm{"V8", false}
	car.turnOnTurbo()

	var car2 sports = &bwm{"V2", false}
	car2.turnOnTurbo()
	fmt.Println(car, car2)

}
