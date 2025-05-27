package main

import "fmt"

// Анонимная функция как аргумент функции
func action(n1 int, n2 int, operation func(int, int) int) {

	result := operation(n1, n2)
	fmt.Println(result)
}

// Анонимная функция как результат функции
func selectFn(n int) func(int, int) int {
	if n == 1 {
		return func(x int, y int) int { return x + y }
	} else if n == 2 {
		return func(x int, y int) int { return x - y }
	} else {
		return func(x int, y int) int { return x * y }
	}
}

// Рекурсивные функции
// Факиториал числа
func factorial(n uint) uint {

	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

//Фибоначи
func fibbonachi(n uint) uint {

	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}

func main() {

	action(10, 25, func(x int, y int) int { return x + y }) // 35
	action(5, 6, func(x int, y int) int { return x * y })   // 30

	f := selectFn(1)
	fmt.Println(f(2, 3)) // 5
	fmt.Println(f(4, 5)) // 9
	fmt.Println(f(7, 6)) // 13

	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(6)) // 720

	fmt.Println(fibbonachi(4)) // 3
	fmt.Println(fibbonachi(5)) // 5
	fmt.Println(fibbonachi(6)) // 8
}
