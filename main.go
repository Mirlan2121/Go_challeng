package main

import "fmt"

type Vehicle interface {
	move()
}

type Car struct{ model string }
type Aircraft struct{ model string }

func (c Car) move() {
	fmt.Println(c.model, "едет")
}
func (a Aircraft) move() {
	fmt.Println(a.model, "летит")
}

func main() {

	tesla := Car{"Tesla"}
	volvo := Car{"Volvo"}
	boeing := Aircraft{"Boeing"}

	vehicles := [...]Vehicle{tesla, volvo, boeing}
	for _, vehicle := range vehicles {
		vehicle.move()
	}
}
