// Чтение с консоли
package main

import (
	"fmt"
)

func main() {
	var name string
	var age int
	fmt.Print("Введите имя и возраст: ")
	fmt.Scan(&name, &age)
	fmt.Println(name, age)

	// альтернативный вариант
	//fmt.Println("Введите имя и возраст:")
	//fmt.Scanf("%s %d", &name, &age)
	//fmt.Println(name, age)
}
