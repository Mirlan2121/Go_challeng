package main

import "fmt"

// Структура контакт
type contanct struct {
	email string
	phone string
}

// Сама структура персон
type person struct {
	name string // Поля имени
	age  int    // Поля возроста
	contanct
}

// Хранения ссылки на структуру того же типа
type node struct {
	value int
	next  *node
}

func printNodeValue(n *node) {
	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}

// Вложенные структуры
func main() {

	var tom = person{
		name: "Tom",
		age:  24,
		contanct: contanct{
			email: "tom@gmail.com",
			phone: "+1234567899",
		},
	}
	tom.email = "supertom@gmail.com"

	fmt.Println(tom.email) // supertom@gmail.com
	fmt.Println(tom.phone) // +1234567899
	fmt.Println()
	// В данном случае структура person имеет поле contactInfo, которое представляет другую структуру contact.

	/*
		Поле contact в структуре person фактические эквивалентно свойству contact contact, то есть свойство называется contact и
		представляет тип contact. Это позволяет нам сократить путь к полям вложенной структуры. Например, мы можем написать
		tom.email, а не tom.contact.email. Хотя можно использовать и второй вариант.
	*/
	////////////////////////////////

	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
	/*
		Здесь определена структура node, которая представляет типичный узел односвязного списка. Она хранит значение в поле value
		и ссылку на следующий узел через указатель next.
		В функции main создаются три связанных структуры, и с помощью цикла for и вспомогательного указателя current выводятся их
		значения.
	*/
}
