// Стандартные потоки ввода-вывода и io.Copy
package main

import (
	"fmt"
	"io"
	"os"
)

/*
	И при выводе из файла текстовой информации на консоль гораздо проще передать данные из файлового потока в os.Stdout,
	через выводить данные отдельными порциями:
*/

//	func main() {
//		file, err := os.Open("hello.txt")
//		if err != nil {
//			fmt.Println(err)
//			os.Exit(1)
//		}
//		defer file.Close()
//
//		io.Copy(os.Stdout, file)
//	}

// В качестве io.Reader можно использовать свои кастомные объекты, которые реализуют данный интерфейс. Например:
type phoneReader string

func (p phoneReader) Read(bs []byte) (int, error) {
	count := 0
	for i := 0; i < len(p); i++ {
		if p[i] >= '0' && p[i] <= '9' {
			bs[count] = p[i]
			count++
		}
	}
	return count, io.EOF
}

func main() {
	phone1 := phoneReader("+1(234)567 90-10")
	io.Copy(os.Stdout, phone1)
	fmt.Println()
}

// В данном случае в качестве интерфейса io.Reader передается объект phoneReader, который считывает цифровые символы
// из номера телефона.
