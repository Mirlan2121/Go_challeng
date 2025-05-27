package main

import "fmt"

// Defer
//Оператор defer позволяет выполнить определенную функцию в
//конце программы, при этом не важно, где в реальности вызывается эта функция.

func main() {
	defer finish()
	fmt.Println("Program has been started")
	fmt.Println("Program is working")
}

func finish() {
	fmt.Println("Program has been finished")
}
