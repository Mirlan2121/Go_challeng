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
