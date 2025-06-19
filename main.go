package main

import "fmt"

func main() {

	intCh := make(chan int, 3)
	intCh <- 10
	intCh <- 3
	close(intCh)
	for i := 0; i < cap(intCh); i++ {
		if val, opened := <-intCh; opened {
			fmt.Println(val)
		} else {
			fmt.Println("Channel closed!")
		}
	}
}
