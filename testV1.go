package main

import (
	"fmt"
)

func main() {
	add := []float64{313, 3312, 14124, 13212, 4114, 24, 88}
	a := summa(add)
	rezult(a, add)
	minimem(add)
	maximal(add)

}

func summa(slis []float64) float64 {
	var sum float64 = slis[0]
	for _, val := range slis {
		sum += val
	}
	fmt.Println("Сумма среза:", sum)
	return sum
}

func rezult(sum float64, srez []float64) (r float64) {
	r = float64(sum) / float64(len(srez))
	fmt.Println("Среднее значение среза:", r)
	return r
}

func minimem(add []float64) float64 {
	var min float64 = add[0]
	for _, val := range add {
		if val < min {
			min = val
		}
	}
	fmt.Println("минимальное", min)
	return min
}

func maximal(srez []float64) float64 {
	var max float64 = srez[0]
	for _, val := range srez {
		if val > max {
			max = val
		}
	}
	fmt.Println("Максимальное число в срезе", max)
	return max
}
