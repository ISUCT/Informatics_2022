package main

//variant 12

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func CalculationX(a float64, b float64, c float64) float64 {
	return (math.Pow(math.E, math.Sin(a)) + math.Log(math.Abs(math.Cos(b)))) / (3 * c)
}

func CalculationY(a float64, с float64) float64 {
	return math.Sqrt(math.Abs(math.Tan(a) + math.Pow(math.Sin(с), float64(1)/3)))
}

func main() {
	var r *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

	fmt.Println("Results calculation x:")
	for i := 1; i < 10; i++ {
		var (
			a float64 = float64(r.Intn(100))
			b float64 = float64(r.Intn(100))
			c float64 = float64(r.Intn(100))
		)
		fmt.Println("a: ", a, "\tb: ", b, "\tc: ", c, "\tresult: ", CalculationX(a, b, c))
	}
	fmt.Println("Results calculation y:")
	for i := 1; i < 10; i++ {
		var (
			a float64 = float64(r.Intn(100))
			c float64 = float64(r.Intn(100))
		)
		fmt.Println("a: ", a, "\tc: ", c, "\tresult: ", CalculationY(a, c))
	}
}
