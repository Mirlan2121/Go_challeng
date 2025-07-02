// Форматируемый ввод
package main

import (
	"fmt"
	"io"
	"os"
)

type person struct {
	name   string
	age    int32
	weight float64
}

func main() {
	filename := "people.dat"
	writeData(filename)
	readData(filename)
}

func writeData(filename string) {
	// начальные данные
	var people = []person{
		{"Tom", 24, 68.5},
		{"Bob", 25, 64.2},
		{"Sam", 27, 73.6},
	}

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for _, p := range people {
		fmt.Fprintf(file, "%s %d %.2f\n", p.name, p.age, p.weight)
	}
}
func readData(filename string) {

	var name string
	var age int
	var weight float64

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer file.Close()

	for {
		_, err = fmt.Fscanf(file, "%s %d %f\n", &name, &age, &weight)
		if err != nil {
			if err == io.EOF {
				break
			} else {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		fmt.Printf("%-8s %-8d %-8.2f\n", name, age, weight)
	}
}
