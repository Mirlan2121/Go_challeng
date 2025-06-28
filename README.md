Для обработки ошибок язык Go предоставляет специальный интерфейс - error, который содержит один метод Error с возвращаемым типом string.

type error interface{
 
    Error() string
}
Значение по умолчанию для типа error — nil. Тип ошибки определяется следующим образом:

Любой тип, реализующий интерфейс error, может представлять тип ошибки. Например:


package main 
import "fmt"
 
type param_error struct{}
 
// реализуем интерфейс error для типа param_error
func (error_object *param_error) Error() string{ 
 
    return "Invalid parameter"
}
 
func main(){
 
    obj := param_error{}
    fmt.Println(obj.Error())    // Invalid parameter
 
    fmt.Println(¶m_error{})   // Invalid parameter
}
Здесь определена пустая структура param_error, которая реализует интерфейс error. В реализации метода Error возвращаем строку "Invalid parameter".

Далее в программе мы можем создать объект этой структуры и вызывать ее реализованный метод Error


obj := param_error{}
fmt.Println(obj.Error())    // Invalid parameter
Либо использовать указатель на безымянный объект этой структуры:


fmt.Println(¶m_error{})   // Invalid parameter
Посмотрим применение на примере:

package main 
import "fmt"
 
type param_error struct{}
 
func (error_object *param_error) Error() string{ 
 
    return "Invalid parameter"
}
 
func factorial(n int) (int, error){
 
    // если некорреткный параметр
    if(n < 0){
        return 0, ¶m_error{}
    }
     
    result := 1
    for i := 1; i <= n; i +=1{
        result *= i
    }
    return result, nil
}
 
func main(){
 
    // Корректный параметр
    fact, err := factorial(5)
    fmt.Println("Factorial:", fact)     // Factorial: 120
    fmt.Println("Error:", err)  // Error: <nil>
 
    // Некорректный параметр
    fact, err = factorial(-5)
    fmt.Println("Factorial:", fact)     // Factorial: 0
    fmt.Println("Error:", err)  // Error: Invalid parameter
}
В данном случае определяем функцию factorial, которая возвращает факториал числа. Причем в реальности возвращается два значения - целое число и ошибку в виде объекта error:


func factorial(n int) (int, error){
Если n < 1, возвращает 0 и указатель на объект param_error, сигнализируя о некорректном параметре:


if(n < 0){
    return 0, ¶m_error{}
}
Иначе вычисляет факториал n в цикле и возвращает результат и nil как ошибку:


result := 1
for i := 1; i <= n; i +=1{
    result *= i
}
return result, nil
Функция main выполняет два вызова factorial. В первом случае передается валидное число 5:


fact, err := factorial(5)
fmt.Println("Factorial:", fact)     // Factorial: 120
fmt.Println("Error:", err)  // Error: <nil>
Во втором случае передается некорректное число -5, поэтому переменная err будет представлять информацию об ошибке в виде объекта param_error:

fact, err := factorial(-5)
fmt.Println("Factorial:", fact)     // Factorial: 0
fmt.Println("Error:", err)  // Error: Invalid parameter
Пакет error
В стандартной библиотеке языка Go имеется встроенный пакет errors, который реализует различные функции для управления ошибками. Он имеет функцию под названием New(), 
которая возвращает объекта интерфейса error и устанавливает его сообщение об ошибке. Например, используем этот метод в нашей функции факториала:


package main 
import (
    "fmt"
    "errors"
)
 
func factorial(n int) (int, error){
 
    if(n < 0){
        return 0, errors.New("недопустимое число: должно быть положительным")
    }
     
    result := 1
    for i := 1; i <= n; i +=1{
        result *= i
    }
    return result, nil  // Возвращаем nil в качестве ошибки, если все хорошо
}
 
func main(){
 
    // Корректный параметр
    fact, err := factorial(5)
    fmt.Println("Factorial:", fact)     // Factorial: 120
    fmt.Println("Error:", err)  // Error: <nil>
 
    // Некорректный параметр
    fact, err = factorial(-5)
    fmt.Println("Factorial:", fact)     // Factorial: 0
    fmt.Println("Error:", err)  // Error: недопустимое число: должно быть положительным
}
