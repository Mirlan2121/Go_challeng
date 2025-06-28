package main

import (
	"errors"
	"fmt"
)

func factorial(n int) (int, error) {

	if n < 0 {
		return 0, errors.New("недопустимое число: должно быть положительным")
	}

	result := 1
	for i := 1; i <= n; i += 1 {
		result *= i
	}
	return result, nil // Возвращаем nil в качестве ошибки, если все хорошо
}

func main() {

	// Корректный параметр
	fact, err := factorial(5)
	fmt.Println("Factorial:", fact) // Factorial: 120
	fmt.Println("Error:", err)      // Error: <nil>

	// Некорректный параметр
	fact, err = factorial(-5)
	fmt.Println("Factorial:", fact) // Factorial: 0
	fmt.Println("Error:", err)      // Error: недопустимое число: должно быть положительным
}
