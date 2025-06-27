package main

import (
	"fmt"
	"sync"
)

type Map_struct struct {
	values map[int]int
	sync.RWMutex
}

var wg sync.WaitGroup

// изменение данных
func (map_data *Map_struct) update_map(keyval int) {

	map_data.Lock() // блокировка на запись

	map_data.values[keyval] = keyval * 10 // изменяем общий ресурс

	map_data.Unlock() // сбрасываем блокировку на запись

	wg.Done() // уведомляем, что горутина завершена
}

// чтение данных
func (map_data *Map_struct) read_map() {

	map_data.RLock() // устанавливаем блокировку на чтение

	val := map_data.values // считываем данные

	map_data.RUnlock() // сбрасываем блокировку на чтение

	fmt.Println(val)
	wg.Done() // уведомляем, что горутина завершена
}

func main() {

	my_map := Map_struct{values: make(map[int]int)}

	wg.Add(2 * 5) // по пять горутин update_map и read_map

	for n := 1; n <= 5; n++ {

		go my_map.update_map(n)

		go my_map.read_map()

	}
	wg.Wait() // ждем завершения всех горутин
}

// RLock(): используется для получения блокировки на чтение. Несколько читателей могут получить эту блокировку в одной программе. Блокировка на чтение не привязан к определенному потоку.

// Unlock(): освобождает полученную одну блокировку на чтение.

// Lock(): блокирует мьютекс для записи. Доступен только для одного потока за раз. Если другой поток получает блокировку для целей чтения/записи, Lock блокируется до тех пор, пока
// блокировка не станет доступной для получения.

// Unlock(): снимает блокировку на запись. При выполнении Unlock без получения блокировки мы получим ошибку во время выполнения.
