// Функции Fprint и Fprintln
package main

import "fmt"

type person struct {
	name   string
	age    int32
	weight float64
}

func main() {
	tom := person{
		name:   "Tom",
		age:    24,
		weight: 68.5,
	}
	fmt.Printf("%-10s %-10d %-10.3f\n",
		tom.name, tom.age, tom.weight)
	fmt.Print("Hello ")
	fmt.Println("cold!")
}
