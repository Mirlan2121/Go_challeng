package main

import "fmt"

func main() {

	str := "Я и есть шторм"
	fmt.Println(str)

	// Используем возвращаемое значение hello()
	greeting := hello()
	fmt.Println("Результат:", greeting)

	// Используем возвращаемые значения Calculator()
	sum, diff, prod, quot := Calculator()
	fmt.Printf("Calculator: +=%d, -=%d, *=%d, /=%d\n", sum, diff, prod, quot)

	// ...

	sum = Add(1, 2, 3, 5, 1)
	fmt.Println("Сумма чисел:", sum)

	powerResult := Power(1.3, 2)
	fmt.Println("1.3^2 =", powerResult)
}

func hello() string {
	num1 := 10
	num2 := 20
	if num1 < num2 {
		return "Правильно"
	}
	return ""
}

func Calculator() (int, int, int, int) {
	num1 := 100
	num2 := 100
	return num1 + num2, num1 - num2, num1 * num2, num1 / num2
}

func Add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Power(n float64, exp int) float64 {
	if exp == 0 {
		return 1.0
	}

	result := 1.0
	absExp := exp
	if exp < 0 {
		absExp = -exp
	}

	for i := 0; i < absExp; i++ {
		result *= n
	}

	if exp < 0 {
		return 1 / result
	}
	return result
}
