package main

import "fmt"

func main() {
	var str string

	str = "Я и есть шторм"
	println(str)

	hello()
	Calculator()

	add(1, 2, 3.4, 5.6, 1.2)
}

// Условия if else
func hello() {
	num1 := 10
	num2 := 20
	if num1 < num2 {
		fmt.Println("Неправильно")
	} else if num1 > num2 {
		fmt.Println("Правильно")
	} else {
		fmt.Println("Равно")
	}
}

// Тенарные операторы
func Calculator() {
	num1 := 100
	num2 := 100
	sum := num1 + num2
	sum1 := num1 * num2
	sum2 := num1 / num2
	sum3 := num1 - num2
	num1++
	num2--

	fmt.Println(sum, sum1, sum2, sum3)
}

//Функции и их параметры
func add(x, y int, a, b, c float32) { // x, y, a, b, c внутернние значения параметров
	var z = x + y
	var d = a + b + c

	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}
