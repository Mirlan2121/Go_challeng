package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0 //  общий ресурс
var wg sync.WaitGroup

func work(number int) {
	counter = 0 // сбрасываем общий ресурс

	for k := 1; k <= 5; k++ {
		counter += 1                       // изменяем общий ресурс
		time.Sleep(100 * time.Millisecond) // задержка для наглядности
		fmt.Println("Goroutine", number, "-", counter)
	}
	wg.Done() // сигнализируем, что горутина завершила работу
}

func main() {

	goroutines_count := 4 // количество запускаемых горутин

	wg.Add(goroutines_count)

	// запускаем горутины
	for i := 1; i <= goroutines_count; i++ {
		go work(i)
	}
	// ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("The End")
}
