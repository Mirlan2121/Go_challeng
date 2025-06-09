package main

import "fmt"

type person struct {
	name string
	age  int
}

// Методы указателей
func (p *person) updateAge(newAge int) {
	(*p).age = newAge
}

func main() {
	var tom = person{name: "Tom", age: 24}
	fmt.Println("before", tom.age)
	tom.updateAge(33)
	fmt.Println("after", tom.age)
}

/*
	Для структуры person определен метод updateAge, который принимает параметр newAge и изменяет значение поля age у структуры. То есть при вызове этого метода мы ожидаем, что
	возраст человека изменится. Однако консольный вывод нам показывает, что значение поля age не изменяется:
	before 24
	after 24

	Однако такое поведение может быть нежелальтельным. Что если мы всетаки хотим таким образом изменять состояние структуры? В этом случае необходимо использовать указатели на
	структуры:
*/
