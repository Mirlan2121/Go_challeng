package main

import "fmt"

// Сама структура персон
type person struct {
	name string // Поля имени
	age  int    // Поля возроста

}

func main() {

	// Первый вариант заполнения
	tom := person{name: "Tom", age: 22}
	var tomPointer *person = &tom

	// Второй вариант заполнения
	var tom1 *person = &person{name: "Tom", age: 24}
	var bob *person = new(person)
	bob.name = "Bob"
	bob.age = 34

	fmt.Println(tom1.name, tom1.age)
	fmt.Println(bob.name, bob.age)

	tomPointer.age = 29
	fmt.Println(tom.age)
	fmt.Println()

	(*tomPointer).age = 32
	fmt.Println(tom.age)
	fmt.Println()
}
