package main

/*
* Syntax Go. Homework 1.1
* Olesya Stetsyak, dated Apr 19, 2020
 */

import (
	"fmt"
)

func main() {
	var doll, euro = 74.1, 81.9
	fmt.Println("Guess the rubles: ")
	var rubles float64
	fmt.Scanln(&rubles)
	var converterDoll = rubles / doll
	var converterEuro = rubles / euro
	fmt.Printf("doll: %.2f, euro: %.2f", converterDoll, converterEuro)
}
