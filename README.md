Оператор panic представляет еще один способ обработки ошибок. Он позволяет сгенерировать ошибку и выйти из программы. 
Он используется для проверки непредвиденных ошибок во всей программе и чем-то похож на исключения в других языках программирования. 
Всякий раз, когда в любой функции вызывается оператор panic, он останавливает выполнение этой функции и завершает выполнение программы. 
Но если в функции, где произошла ошибка, вызывваются функции с операторм defer, то перед завершением программы выполняются эти функции.

В качестве параметра оператор panic принимает произвольное значение (формально интерфейс interface {}, которому в реальности соответствует произвольное значение):


panic (interface {})
Пример применения:


package main
import "fmt"
 
func main() {
    fmt.Println(divide(15, 5))
    fmt.Println(divide(4, 0))
    fmt.Println("Program has been finished")
}
func divide(x, y float64) float64{
    if y == 0{ 
        panic("Division by zero!")
    }
    return x / y
}
В данном случае в функции divide, если второй параметр равен 0, то осуществляется вызов panic("Division by zero!").

В функции main в вызове fmt.Println(divide(4, 0)) будет выполняться оператор panic, поскольку второй параметр функции divide равен 0. 
И в этом случае все последующие операции, которые идут после этого вызова, например, в данном случае это вызов fmt.Println("Program has been finished"), 
не будут выполняться. В этом случае мы получим следующий консольный вывод:


panic: Division by zero!
И в конце вывода будет идти диагностическая информация о том, где возникла ошибка.

В некоторых случаях мы можем выполнить некоторую очистку или другую необходимую операцию, прежде чем программа завершится. 
Это можно сделать, используя функцию с оператором defer, которая выполняется перед завершением программы:


package main
import "fmt"
 
func divide_defer(){
 
    fmt.Println("divide_defer executed")
}
 
func main_defer(){
 
    fmt.Println("main_defer executed")
}
 
 
func divide(x, y float64) float64{
    defer divide_defer()
    if y == 0{ 
        panic("Division by zero!")
    }
    return x / y
}
 
func main() {
    defer main_defer()
    fmt.Println(divide(4, 0))
    fmt.Println("Program has been finished")
}
Консольный вывод:

divide_defer executed
main_defer executed
panic: Division by zero!
