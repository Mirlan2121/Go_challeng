package main

import "fmt"

func try_catch() {

	if r := recover(); r != nil {
		fmt.Println("Error:", r)
	}
}

func divide(x, y float64) float64 {
	defer try_catch()
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}

func main() {
	fmt.Println(divide(4, 0))
	fmt.Println("Program has been finished")
}
