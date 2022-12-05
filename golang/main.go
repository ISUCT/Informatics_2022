package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("ZADA")
	fmt.Println(zadA(1.25, 3.25, 0.4))
	fmt.Println("ZADB")
	fmt.Println(zadB([]float64{1.84, 2.71, 3.81, 4.56, 5.62}))
}
func zadA(xn float64, xk float64, dx float64) []float64 {
	var x float64 = xn
	var massivA = make([]float64, 0)
	for ; x <= xk; x += dx {
		massivA = append(massivA, form(x))
	}
	return massivA
}
func zadB(massivB []float64) []float64 {
	var i int8 = 0
	for ; i <= 4; i += 1 {
		massivB[i] = form(massivB[i])
	}
	return massivB
}
func form(y float64) float64 {
	return (math.Sqrt(math.Sqrt(math.Abs(y*y-2.5))) + math.Cbrt(math.Log10(y*y)))
}
