package main

/*
* Syntax Go. Homework 7.3
* Olesya Stetsyak, dated 18 May, 2020
 */
import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
	"time"
)

func main() {
	go waitExit() // горутина ожидания выхода
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn)
	}
}

func waitExit() { // функция завершения работы сервера
	scanner := bufio.NewScanner(os.Stdin) // создаём новый объект сканер
	for scanner.Scan() {                  // считываем из консоли
		if scanner.Text() == "exit" { // если, то что ввели в консоль слово "exit", завершаем работу сервера
			os.Exit(0)
		}
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n\r"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
