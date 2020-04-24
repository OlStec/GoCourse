package main

/*
* Syntax Go. Homework 2.1
* Olesya Stetsyak, dated Apr 23, 2020
 */

import (
	"fmt"
)

func main() {
	res := Even(2)
	fmt.Printf("%d even: %t\n", 2, res)
}

func Even(num int) bool {
	return num%2 == 0
}
