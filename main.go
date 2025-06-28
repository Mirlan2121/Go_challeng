package main

import "fmt"

func factorial(n int) (int, interface{}) {

	// если переданное число меньше 0
	if n < 0 {
		return 0, "Недопустимое число. Должно быть больше 0"
	}

	result := 1
	for i := 1; i <= n; i += 1 {
		result *= i
	}
	return result, nil
}

func main() {

	// Корректный параметр
	fact, err := factorial(5)
	fmt.Println("Factorial:", fact) // Factorial: 120
	fmt.Println("Error:", err)      // Error: <nil>

	// Некорректный параметр
	fact, err = factorial(-5)
	fmt.Println("Factorial:", fact) // Factorial: 0
	fmt.Println("Error:", err)      // Error: Недопустимое число. Должно быть больше 0
}

//func main() {
//
//	file, err := os.Open("./main.go")
//	fmt.Println("file :", file)
//	fmt.Println("error:", err)
// }
