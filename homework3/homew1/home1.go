package main

/*
* Syntax Go. Homework 3.1&2
* Olesya Stetsyak, dated Apr 28, 2020
 */

import (
	"fmt"
)

// Vehicle автомобили и их характеристи
type Vehicle struct {
	brand           string
	year            int
	trunkVolume     int
	isEngineRunning bool
	areWindowsOpen  bool
	trunkFull       int
	color           string
}

//SetColor - устанавливаев значение в Vehicle
func (v *Vehicle) SetColor(color string) {
	v.color = color
}

//GetColor - устанавливаев значение
func (v *Vehicle) GetColor() string {
	return v.color
}

func main() {
	truck := Vehicle{"volvo", 2017, 10000, false, false, 10, "red"}
	car := Vehicle{"Mersedes", 2017, 300, false, false, 0, "yellow"}
	car.trunkFull = 70 // Заполнили на 70% бак автомобиля

	car.SetColor("blue")
	fmt.Println(car.GetColor())

	fmt.Println(truck)
	fmt.Println(car)
	fmt.Println(truck.year == car.year) // Сравнили годы выпуска транспорта
}
