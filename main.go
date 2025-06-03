package main

import "fmt"

// Указатели и функции

func changeValue(x int) {
	x = x * x
}

// Указатели как параметры функции
func changeValue2(x *int) {
	*x = (*x) * (*x)
}

// Указатель как результат функции
func createPointer(x int) *int {
	p := new(int)
	*p = x
	return p
}

func main() {

	fmt.Println("Первый вариант кода без указателей")
	d := 5
	fmt.Println("d before: ", d) // 5
	changeValue(d)               // Изменением значение
	fmt.Println("d after: ", d)  // 5 - значение неизменилось
	fmt.Println()
	fmt.Println()

	/*
			Функция changeValue изменяет значение параметра, возводя его в квадрат. Но после вызова этой функции мы видим, что значение
			переменной d, которая передается в changeValue, не изменилось. Ведь	функция получает копию данной переменной и
		    работает с ней независимо от оригинальной переменной d. Поэтому d никак не изменяется.
	*/
	fmt.Println("Второй вариант кода с указателим")
	d2 := 5
	fmt.Println("d before:", d2) // 5
	changeValue2(&d2)            // изменяем значение
	fmt.Println("d after:", d2)  // 25 - значение изменилось!
	fmt.Println()
	fmt.Println()
	/*
	   Теперь функция changeValue принимает в качестве параметра указатель на объект типа int. При вызове функции changeValue в нее
	   передается адрес переменной d (changeValue(&d)). И после ее выполнения мы видим, что значение переменной d изменилось.
	*/

	// В данном случае функция createPointer возвращает указатель на объект int.
	fmt.Println("В данном случае функция createPointer возвращает указатель на объект int.")
	p1 := createPointer(7)
	fmt.Println("p1:", *p1) // p1: 7
	p2 := createPointer(10)
	fmt.Println("p2:", *p2) // p2: 10
	p3 := createPointer(28)
	fmt.Println("p3:", *p3) // p3: 28
}
