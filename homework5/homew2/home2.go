package main

/*
* Syntax Go. Homework 5.2
* Olesya Stetsyak, dated 12 May, 2020
 */

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("fileread.go") // создаём фаил go, который учавствует в коде, что бы комп. не ругался
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	file, err = os.Open("fileread.go")
	if err != nil {
		log.Fatal(err) // как я поняла правильно всегда возвращать ошибку, а не просто return
	}
	defer file.Close()

	// getting size of file
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err) // возвращаем ошибку
	}

	// reading file
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err) // возвращаем ошибку
	}
}
