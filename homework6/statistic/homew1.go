package statistic

/*
* Syntax Go. Homework 6.1
* Olesya Stetsyak, dated 13 May, 2020
 */

//Average - находит среднее значение чиисел
func Average(xs []float64) float64 {
	total := float64(0)
	for _, x := range xs {
		total += x
	}
	return total / float64(len(xs))
}

//Amount - находит сумму чисел
func Amount(xs []float64) float64 {
	var sum float64
	for _, elem := range xs {
		sum += elem
	}
	return sum
}
