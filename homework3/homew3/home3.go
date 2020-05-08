package main

/*
* Syntax Go. Homework 3.3
* Olesya Stetsyak, dated Apr 28, 2020
 */

import (
	"fmt"
)

// Queue Заполняемая очередь
var Queue []string

func insertQueue(str ...string) {
	Queue = append(Queue, str...) // Добавление в очередь
}

func dequeueQueue() string {
	if isEmptyQueue() {
		return ""
	}
	firstItem := Queue[0] // Первый элемент
	Queue = Queue[1:]     // Удаление из очереди
	return firstItem
}

func isEmptyQueue() bool {
	return len(Queue) == 0
}

func main() {
	fmt.Println("Пустая очередь", isEmptyQueue())
	insertQueue("Hello ")
	insertQueue("world!")
	fmt.Println(Queue)
}
