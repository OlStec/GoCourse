package main

/*
* Syntax Go. Homework 2.4
* Olesya Stetsyak, dated Apr 24, 2020
 */

import (
	"fmt"
)

func seive(n int) {
	seiveArr := make([]int, n)
	for i := 2; i*i < n; i++ {
		if seiveArr[i] == 0 {
			for k := i * i; k < n; k += i {
				seiveArr[k] = 1
			}
		}
	}
	fmt.Println("Простые числа:\n")

	for i := 2; i < n; i++ {
		if seiveArr[i] == 0 {
			fmt.Printf("%3d", i)
		}
	}
}

const n1 = 100

func main() {
	seive(n1)
}
