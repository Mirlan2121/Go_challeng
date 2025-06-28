package main

import "fmt"

func divide_defer() {

	fmt.Println("divide_defer executed")
}

func main_defer() {

	fmt.Println("main_defer executed")
}

func divide(x, y float64) float64 {
	defer divide_defer()
	if y == 0 {
		panic("Division by zero!")
	}
	return x / y
}

func main() {
	defer main_defer()
	fmt.Println(divide(4, 0))
	fmt.Println("Program has been finished")
}

//Оператор defer помещает вызов переданной ему функции в стек. Когда окружающая функция возвращает управление, вызываются все
//функции, хранящиеся в стеке. Функция, хранящаяся в конце стека, выполняется первой. Значения, переданные в аргументах функции
//defer, вычисляются, когда компилятор находит оператор defer.

//В функциях defer могут выполняться как операции чтения, так и записи.

//Возвращаемые значения могут быть изменены в функциях defer.

//Оператор defer обеспечивает выполнение отложенных функций даже в случае ошибки или вызова функции panic, помогая очищать
//ресурсы.

// Пример
//package main
//import "fmt"
//
//func main() {
//    defer finish()
//    fmt.Println("Program has been started")
//    fmt.Println("Program is working")
//}
//
//func finish(){
//    fmt.Println("Program has been finished")
//}
