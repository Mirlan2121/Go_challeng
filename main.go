package main

import "fmt"

// Интерфейс
type Vehicle interface {
	move()
}

// Функция движения транспорта
func drive(vehicle Vehicle) {
	vehicle.move()
}

// Структура машины и самолета
type Car struct{}
type Aircraft struct{}

// Функции движения транспорта
func (c Car) move() {
	fmt.Println("Автомобиль едет")
}
func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func main() {
	tesla := Car{}
	boing := Aircraft{}
	drive(tesla)
	drive(boing)

}
