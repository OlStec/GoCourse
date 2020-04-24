package main

/*
* Syntax Go. Homework 1.2
* Olesya Stetsyak, dated Apr 19, 2020
 */

import (
	"fmt"
)

func main() {
	var trapBaseLen1, trapBaseLen2, trapeBaseH float64
	fmt.Println("Введите длинну №1 основания трапеции:")
	fmt.Scanln(&trapBaseLen1)
	fmt.Println("Введите длинну №2 основания трапеции:")
	fmt.Scanln(&trapBaseLen2)
	fmt.Println("Введите высоту трапеции:")
	fmt.Scanln(&trapeBaseH)

	var trapeArea float64 = ((trapBaseLen1 + trapBaseLen2) / 2) * trapeBaseH
	fmt.Printf("Area: %.2f", trapeArea)
}
