package main

/*
* Syntax Go. Homework 7.2
* Olesya Stetsyak, dated 17 May, 2020
 */
import (
	"fmt"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	// генерация
	go func() {
		for x := 0; x < 100; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	// возведение в квадрат
	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()

	// печать
	for square := range squares {
		fmt.Println(square)
	}
}
