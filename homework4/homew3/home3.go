package main

/*
* Syntax Go. Homework 4.31
* Olesya Stetsyak, dated 7 May, 2020
 */
import "fmt"

func main() {
	var firstArg, secondArg float64
	var operator string
	var help string

	fmt.Println("Введите слово help, чтобы вызвать справку или нажмите Enter для продоления:")
	fmt.Scanln(&help)
	helpCalculate(help)

	fmt.Println("Введите первый аргумент:")
	fmt.Scanln(&firstArg)

	fmt.Println("Введите оператор:")
	fmt.Scanln(&operator)

	fmt.Println("Введите второй аргумент:")
	fmt.Scanln(&secondArg)

	result, err := calculate(firstArg, operator, secondArg)
	if err != nil {
		fmt.Printf("Пройзошла ошибка: %s", err.Error())
		return
	}

	fmt.Println("Результат:", result)
}
