package main

import "fmt"

func main() {

	intCh := make(chan int)

	go square(intCh)                   // square ожидает получения через канал
	intCh <- 4                         // отправляем в канал число
	fmt.Println("result := ", <-intCh) // получаем из канала результат
	fmt.Println("The End")
}

// функция возведения в квадрат
func square(ch chan int) {

	num := <-ch // получаем из канала число
	fmt.Println("num := ", num)
	ch <- num * num // обратно отправляем квадрат числа
}
