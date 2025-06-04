package main

import "fmt"

type mile uint
type kilometer uint
type BinaryOp func(int, int) int

func action1(n1 int, n2 int, op BinaryOp) {
	result := op(n1, n2)
	fmt.Println(result)
}

func add1(x int, y int) int {
	return x + y
}

/*
   В данном случае определяется именованный тип mile, который основывается на типе uint. По сути mile представляет тип uint и
   работа с ним будет производиться также, как и с типом uint. Однако в то же время фактически это новый тип.
*/

func distanceToEnemy(distance mile) {
	fmt.Println("расстояние для противника:")
	fmt.Println(distance, "миль")

}
func action(n1 int, n2 int, op func(int, int) int) {
	result := op(n1, n2)
	fmt.Println(result)
}

func add(x int, y int) int {
	return x + y
}

func main() {
	// Мы можем определять переменные данного типа, работать с ними как с объектами базового типа uint:

	var distance mile = 5
	fmt.Println(distance)
	distance += 5
	fmt.Println(distance)
	fmt.Println()
	fmt.Println()

	/*
	   Но может возникнуть вопрос, а зачем это нужно, зачем определять именнованный тип, если он все равно ведет себя как
	   тип uint? Рассмотрим следующую ситуацию:
	*/

	var distance1 mile = 5
	distanceToEnemy(distance1)
	fmt.Println()

	// var distance1 uint = 5
	// distanceToEnemy(distance1)   // !Ошибка

	// var distance2 kilometer = 5
	// distanceToEnemy(distance2)   // ! ошибка

	/*
		Здесь определена функция action, которая принимает два числа и некоторую другую функцию с типом func(int, int) int -
		то есть функцию, которая принимает два числа и также возвращает число. В функции main определяется переменная myOperation,
		которая как раз представляет функцию типа func(int, int) int, получает ссылку на функцию add и передается в вызов
		action(10, 25, myOperation)
	*/

	/*
	Теперь тип функции func(int, int) int проецируется на именнованный тип BinaryOp, который представляет бинарную операцию над двумя операндами:


	*/
	var myOperation func(int, int) int = add
	action(10, 25, myOperation)
	fmt.Println()
	fmt.Println()

	var myOperation1 BinaryOp = add1
	action1(10, 35, myOperation1)
}
