# 🚀 Типы и Функции в Go

![Go Logo](https://golang.org/lib/godoc/images/go-logo-blue.svg)

Этот репозиторий содержит примеры и объяснения различных типов функций в языке программирования Go.

## 📌 Содержание

1. [Базовые функции](#-базовые-функции)
2. [Функции с несколькими возвращаемыми значениями](#-функции-с-несколькими-возвращаемыми-значениями)
3. [Анонимные функции](#-анонимные-функции)
4. [Замыкания](#-замыкания)
5. [Функции как параметры](#-функции-как-параметры)
6. [Методы](#-методы)

## 🧩 Базовые функции

// Простая функция
func greet(name string) string {
    return "Hello, " + name
}

// Функция без параметров
func version() string {
    return "1.0.0"
}

## 🔄 Функции с несколькими возвращаемыми значениями

// Возврат нескольких значений
func divide(a, b float64) (float64, error) {
    if b == 0.0 {
        return 0.0, fmt.Errorf("cannot divide by zero")
    }
    return a / b, nil
}

// Именованные возвращаемые значения
func calc(a, b int) (sum int, diff int) {
    sum = a + b
    diff = a - b
    return
}

## 🎭 Анонимные функции
func main() {
    // Анонимная функция
    func(msg string) {
        fmt.Println(msg)
    }("Hello from anonymous function!")
}

## 🔗 Замыкания
func counter() func() int {
    i := 0
    return func() int {
        i++
        return i
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
}

## 📤 Функции как параметры
func apply(numbers []int, op func(int) int) []int {
    result := make([]int, len(numbers))
    for i, v := range numbers {
        result[i] = op(v)
    }
    return result
}

func square(n int) int {
    return n * n
}

func main() {
    nums := []int{1, 2, 3}
    squared := apply(nums, square)
    fmt.Println(squared) // [1 4 9]
}

## 🏗 Методы
type Rectangle struct {
    Width, Height float64
}

// Метод для типа Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

func main() {
    rect := Rectangle{10, 5}
    fmt.Println(rect.Area()) // 50
}
