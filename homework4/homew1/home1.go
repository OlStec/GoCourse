package main

import (
	"fmt"
)

/*
* Syntax Go. Homework 4.1
* Olesya Stetsyak, dated 4 May, 2020
 */

//RoomSquare - структура данных о комнатах, где пол с керамогранитом
type RoomSquare struct {
	Side float64
}

//RoomRectangle - структура данных о комнатах, где пол с ковралином
type RoomRectangle struct {
	Lenght float64
	Width  float64
}

//Shape - интерфейс
type Shape interface {
	Area() float64
}

//Area - функция считающая площадь комнат
func (r *RoomSquare) Area() float64 {
	return r.Side * r.Side
}

//Area - функция считающая площадь комнат
func (r *RoomRectangle) Area() float64 {
	return r.Lenght * r.Width
}

//SumAreas - функция считающая сумму площади комнат
func SumAreas(shapes ...Shape) float64 {
	var sum float64
	for _, shape := range shapes {
		sum += shape.Area()
	}
	return sum
}

func main() {
	Room1 := &RoomSquare{2}
	Room4 := &RoomRectangle{3, 6}
	sumAreaRoom := SumAreas(Room1, Room4)
	fmt.Printf("%.2f\n", sumAreaRoom)
}
