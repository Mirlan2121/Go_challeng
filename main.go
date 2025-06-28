package main

import (
	"fmt"
	"os"
)

func open_file() {

	// открываем файл
	file, err := os.Open("./main.go")
	// закрываем файл в конце функции
	defer file.Close()
	// работа с файлом
	fmt.Println("file:", file)
	fmt.Println("error:", err)
}

func main() {

	open_file()
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
