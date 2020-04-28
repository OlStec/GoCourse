package main

/*
* Syntax Go. Homework 4.4
* Olesya Stetsyak, dated Apr 28, 2020
 */

import (
	"fmt"
)

type Person struct {
	Name    string
	Phone   string
	Address string
}
type addressBook struct {
	Person []Person
}

func main() {
	Alex := Person{"Alex Jones", "+ 235 58-80-88", "1801 SW 2nd Ct, Miami, FL 33129, USA"}
	Bob := Person{"Bob Brown", "+ 232 98-00-78", "11301 West Olympic Boulevard Apt. 100 Los Angeles, CA 90064 USA"}
	addressBook := AddressBook{[]Person{Alex, Bob}}
	fmt.Println(addressBook)

}
