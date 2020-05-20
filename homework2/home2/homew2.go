package main

/*
* Syntax Go. Homework 2.2
* Olesya Stetsyak, dated Apr 23, 2020
 */

import (
	"fmt"
)

func main() {
	res1 := result(3)
	fmt.Printf("%d делится нацело на 3 %t\n", 3, res1)

}

func result(num int) bool {
	return num%3 == 0
}
