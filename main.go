package main

import "fmt"

func main() {

	str := "Я и есть шторм"
	fmt.Println(str)

	hello()
	Calculator()

	sum := add(1, 2, 3, 5, 1)
	fmt.Println("Сумма чисел:", sum)

	powerResult := Power(1.3, 2)
	fmt.Println("1.3^2 =", powerResult)
}

func hello() {
	num1 := 10
	num2 := 20
	if num1 < num2 {
		fmt.Println(num1, "<", num2)
	} else if num1 == num2 {
		fmt.Println("Числа равны")
	}
}

func Calculator() {
	num1 := 100
	num2 := 100
	sum := num1 + num2
	sum1 := num1 * num2
	sum2 := num1 / num2
	sum3 := num1 - num2
	fmt.Println("Результаты:", sum, sum1, sum2, sum3)
}

func add(numbers ...int) int {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Power(n float64, exp int) float64 {
	if exp < 0 {
		n = 1 / n
		exp = -exp
	}
	result := 1.0
	for i := 0; i < exp; i++ {
		result *= n
	}
	return result
}
