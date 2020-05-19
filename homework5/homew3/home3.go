package main

/*
* Syntax Go. Homework 5.3
* Olesya Stetsyak, dated 11 May, 2020
 */

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "username"},
		{"Bob", "Brown", "bob"},
		{"Olivia", "Thompson", "olivia"},
		{"Emily", "Walker", "emily"},
	}

	file, err := os.Create("file.csv") // создаём фаил csv
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	write := csv.NewWriter(file) // Cоздаём структуру сsv, разделённую запятыми
	write.WriteAll(records)      // Записываем все строки

	csvfile, err := os.Open("file.csv") // открываем фаил
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	text, err := ioutil.ReadFile("file.csv") // читаем весь фаил
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(text)) // выводим на экран текст в фаиле

}
