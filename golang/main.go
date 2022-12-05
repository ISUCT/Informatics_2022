package main

import (
	"fmt"
	"math"
)

func main() {
	const (
		a float64 = 1.2
		b float64 = 0.48
	)
	fmt.Println(taskA(a, b, 0.7, 2.2, 0.3))
	var users = [5]float64{0.25, 0.36, 0.56, 0.94, 1.28}
	fmt.Println(taskB(a, b, users[:]))
}

func taskA(a, b, xn, xk, dx float64) []float64 {
	var ansA []float64
    for x := xn; x <= xk; x += dx {
        ansA = append(ansA, par(a, b, x))
	}
	return ansA 
}

func taskB(a, b float64, xm []float64) []float64 {
	var ansB []float64
	for _, x := range  xm {
	    ansB = append(ansB, par(a, b, x))
	}
	return ansB 
}

func par(a, b, x float64) float64 {
	y := (math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos(x*b*x) + (math.Exp(-x/2)))
	return y
}