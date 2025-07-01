package main

import (
	"fmt"
	"io"
	"os"
)

// Запись в файл
// Для записи текстовой информации в файл можно применять метод WriteString() объекта os.File, который заносит в файл строку:

/*
func main() {
	text := "Hello Gold!"
	file, err := os.Create("hello.txt")

	if err != nil {
		fmt.Println("Unable to create file:", err)
		os.Exit(1)
	}
	defer file.Close()
	file.WriteString(text)

	fmt.Println("Done.")
}
*/

// В данном случае создается файл hello.txt, в который записывается строка "Hello Gold!".

// Для записи нетекстовой бинарной информации в виде набора байт применяется метод Write() (реализация интерфейса io.Writer):

// func main() {
// 	data := []byte("Hello Bold!")
// 	file, err := os.Create("hello.bin")
// 	if err != nil {
// 		fmt.Println("Unable to create file:", err)
// 		os.Exit(1)
// 	}
// 	defer file.Close()
// 	file.Write(data)
//
// 	fmt.Println("Done.")
// }

// Чтение из файла
// Поскольку тип io.File реализует интерфейс io.Reader, то для чтения из файла мы можем использовать метод Read().
// Этот метод позволяет получить содержимое файла в виде набора байт:

func main() {
	file, err := os.Open("hello.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	data := make([]byte, 64)

	for {
		n, err := file.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		fmt.Print(string(data[:n]))
	}
}

// Для считывания данных определяется срез из 64 байтов. В бесконечном цикле содержимое файла считывается в срез, а
// когда будет достигнут конец файла, то есть мы получим ошибку io.EOF, то произойдет выход из цикла. Ну и поскольку
// анные представляют срез байтов, хотя файл hello.txt хранит текстовую информацию, то для вывода текста на консоль
// преобразуем срез байтов в строку: string(data[:n]).
