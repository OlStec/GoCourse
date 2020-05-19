package statistic

/*
* Syntax Go. Homework 6.1test
* Olesya Stetsyak, dated 13 May, 2020
 */

import "testing"

type testpair struct {
	values  []float64
	average float64
}
type testpairAmount struct {
	valuesAmount []float64
	amount       float64
}

var tests = []testpair{
	{[]float64{1, 2}, 1.5},
	{[]float64{1, 1, 1, 1, 1, 1}, 1},
	{[]float64{-1, 1}, 0},
}

var testsAmount = []testpairAmount{
	{[]float64{1, 2}, 3},
	{[]float64{1, 1, 1, 1, 1, 1}, 6},
	{[]float64{-1, 1}, 0},
}

func TestAverageSet(t *testing.T) {
	for _, pair := range tests {
		v := Average(pair.values)
		if v != pair.average {
			t.Error(
				"For", pair.values,
				"expected", pair.average,
				"got", v,
			)
		}
	}
}
func TestAmountSet(t *testing.T) {
	for _, pair := range testsAmount {
		v := Amount(pair.valuesAmount)
		if v != pair.amount {
			t.Error(
				"For", pair.valuesAmount,
				"expected", pair.amount,
				"got", v,
			)
		}
	}
}
