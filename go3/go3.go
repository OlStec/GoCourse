package main

/*
* Syntax Go. Homework 1.3
* Olesya Stetsyak, dated Apr 19, 2020
 */

import (
	"fmt"
	"math"
)

func main() {

	var y string
	var oneYear float64 = 365
	var years float64 = 2 /* кол-во лет вклада*/
	var deposit, percent float64
	var n float64 = 0
	var sumDeposit float64 = 0
	fmt.Println("Введите сумму вклада:")
	fmt.Scanln(&deposit)
	fmt.Println("Введите процент:")
	fmt.Scanln(&percent)

	for n = 0; n < years; n++ {
		sumDeposit = deposit * math.Pow(1+oneYear*(percent/(100*oneYear)), n)
	}
	if n < 5 && n > 1 {
		y = "года"
	} else if n >= 5 {
		y = "лет"
	}
	fmt.Printf("Сумма вклада с процентами за , %.f, %s, будет составдять: %.2f", n, y, sumDeposit)
}
