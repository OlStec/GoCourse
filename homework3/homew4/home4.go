package main

/*
* Syntax Go. Homework 3.4
* Olesya Stetsyak, dated 2 May, 2020
 */

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

//File - название фаила в кот. будем сохранять
var File string = ("home4add.txt")

//Person - структура данных о человеке
type Person struct {
	ID      int
	Name    string
	Phone   string
	Address string
}

//AddressBook - структура книги с данными о Person
type AddressBook struct {
	Person []Person
}

//MarshalJ - сохраняем Json
func (add AddressBook) MarshalJ() {
	a, err := json.Marshal(add)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(a))
	WriteFile(File, a)
}

//WriteFile - метод, с помощью которого записываем данные в файл "home4add.txt"
func WriteFile(File string, Bookdata []byte) {
	err := ioutil.WriteFile(File, Bookdata, 0644)
	if err != nil {
		fmt.Println(err)
	}
}
func main() {
	Alex := Person{1, "Alex Jones", "+ 235 58-80-88", "1801 SW 2nd Ct, Miami, FL 33129, USA"}
	Bob := Person{2, "Bob Brown", "+ 232 98-00-78", "11301 West Olympic Boulevard Apt. 100 Los Angeles, CA 90064 USA"}
	Olivia := Person{3, "Olivia Smith", "+ 232 96-06-76", "11303 West Olympic Boulevard Apt. 140 Los Angeles, CA 90064 USA"}
	Emily := Person{4, "Emily Walker", "+ 232 98-88-78", "11302 West Olympic Boulevard Apt. 102 Los Angeles, CA 90064 USA"}
	Bookdata := AddressBook{[]Person{Alex, Bob, Olivia, Emily}}
	//fmt.Println(Bookdata)

	Bookdata.MarshalJ()
}
