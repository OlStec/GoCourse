package main

/*
* Syntax Go. Homework 4.32
* Olesya Stetsyak, dated 7 May, 2020
 */
import "fmt"

func calculate(firstAtg float64, operator string, secondArd float64) (float64, error) {
	switch operator {
	case "+":
		return firstAtg + secondArd, nil
	case "-":
		return firstAtg - secondArd, nil
	case "*":
		return firstAtg * secondArd, nil
	case "/":
		if secondArd == 0 {
			return 0, fmt.Errorf("Нельзя делить на ноль")
		}
		return firstAtg / secondArd, nil
	default:
		return 0, fmt.Errorf("Некоректный оператор")
	}
}

func helpCalculate(help string) {
	if help == "help" {
		fmt.Println("После запуска программы введите число или слово help - выводит инструкцию, после чего выходит из программы./n Далее если то арифмитическое действие +, -, *, /, введите его выбрав нужное и далее вводим 2 значение. Получаем результат.")
	}
}
