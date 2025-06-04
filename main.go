package main

import "fmt"

// Сама структура персон
type person struct {
	name string // Поля имени
	age  int    // Поля возроста

}

// Условия проверять на Имя
func addName(p person) string {

	name := "Miki"
	if name == p.name {
		fmt.Println("Имя соответствует")
	} else {
		fmt.Println("Нет такой именни")
	}

	return name
}

// Условия проверять на Возрост
func addAge(p person) {

	age := 14

	if age >= p.age {
		fmt.Println("Возрост соответствует")
	} else {
		fmt.Println("Возрост не соответствует")
	}

}

func main() {

	var tom person = person{name: "Miki", age: 23}
	fmt.Println(tom.name) // Miki
	fmt.Println(tom.age)  // 23
	fmt.Println()

	tom.age = 24                   // Изменения значения
	fmt.Println(tom.name, tom.age) // Miki, 24
	fmt.Println()

	var name person = person{name: "Alex", age: 15}
	fmt.Println(name)

	addName(name)
	addAge(name)
}
