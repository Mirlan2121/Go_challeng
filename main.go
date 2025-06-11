package main

// Импорт пакетов
import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	// Вызов модуля quote и его функции Hello
	message := quote.Hello()
	fmt.Println(message)
}
