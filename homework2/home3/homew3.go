package main

/*
* Syntax Go. Homework 2.3
* Olesya Stetsyak, dated Apr 24, 2020
 */

import (
	"fmt"
)

/* вариант рекурсивного решения
func fibonacci() func() int {
	n1, n := -1, 1
	return func() int {
		n1, n = n, n1+n
		return n
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 100; i++ {
		fmt.Println(f())
	}
}*/

func fib(n int) int {
	if n < 2 {
		return n
	} else {
		return fib(n-1) + fib(n-2)
	}
}

const n = 100

func main() {
	for i := 0; i < n; i++ {
		fmt.Println(fib(i))
	}
}
