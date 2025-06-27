package main

import (
	"fmt"
	"sync"
	"time"
)

var counter int = 0 //  общий ресурс
var wg sync.WaitGroup

func work(number int, mutex *sync.Mutex) {

	mutex.Lock() // блокируем доступ к переменной counter

	counter = 0 // сбрасываем общий ресурс

	for k := 1; k <= 5; k++ {
		counter += 1 // изменяем общий ресурс
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Goroutine", number, "-", counter)
	}

	mutex.Unlock() // деблокируем доступ

	wg.Done() // сигнализируем, что горутина завершила работу
}

func main() {

	var mutex sync.Mutex // определяем мьютекс

	goroutines_count := 4 // количество запускаемых горутин

	wg.Add(goroutines_count)

	// запускаем горутины
	for i := 1; i <= goroutines_count; i++ {
		go work(i, &mutex)
	}
	// ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println("The End")
}
