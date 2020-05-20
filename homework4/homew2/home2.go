package main

/*
* Syntax Go. Homework 4.2
* Olesya Stetsyak, dated 6 May, 2020
 */

import (
	"fmt"
	"sort"
)

//Phones - номера телефонов
type Phones []int

func (p Phones) Len() int {
	return len(p)
}

func (p Phones) Less(i, j int) bool {
	return p[i] < p[j]
}

func (p Phones) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

//PhonesBook  структура описывающая контакты
type PhonesBook struct {
	Name   string
	Phones Phones
}

//Sort - функция для сортировки телефоной книги
func (ph PhonesBook) Sort() {
	sort.Sort(ph.Phones)
}

func main() {
	phonesBook := PhonesBook{"Olivia", Phones{235558000, 235558080, 225367}}
	phonesBook.Sort()
	fmt.Println(phonesBook)
}
