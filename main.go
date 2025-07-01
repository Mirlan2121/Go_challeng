package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("hello.txt") // создаем файл
	if err != nil {                     // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	defer file.Close()       // закрываем файл
	fmt.Println(file.Name()) // hello.txt
}
