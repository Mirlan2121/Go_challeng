package main

import "fmt"

func main() {
	var a = add1(4, 5)  // 9
	var b = add1(20, 6) // 26
	fmt.Println(a)
	fmt.Println(b)
}

func add1(x, y int) int {
	return x + y
}
