package main

import "fmt"

func main() {
	str := "Я и есть шторм"
	fmt.Println(str)

	greeting := hello()
	fmt.Println("Результат:", greeting)

	sum, diff, prod, quot := Calculator()
	fmt.Printf("Calculator: +=%d, -=%d, *=%d, /=%d\n", sum, diff, prod, quot)

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

	if num2 == 0 {
		return 0, 0, 0, 0
	}
	return num1 + num2, num1 - num2, num1 * num2, num1 / num2
}

func Add(numbers ...int) int {
	if len(numbers) == 0 { // Если чисел нет — возвращай 0
		return 0
	}
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	return sum
}

func Power(n float64, exp int) float64 {
	if exp == 0 { // 0^0 = 1 (по соглашению)
		return 1
	}
	if n == 0 {
		return 0 // 0^5 = 0
	}
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

// Новые функции для тестирования
func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Div(a, b int) (int, error) {
	if b == 0 {
		return 0, fmt.Errorf("деление на ноль")
	}
	return a / b, nil
}
