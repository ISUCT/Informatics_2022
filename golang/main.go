package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(zadA())
	fmt.Println(zadB())
}

func zadA() []float64 {
	fmt.Println("ZAD A")
	var i float64 = 1.25
	var massivA []float64
	var count int = 0
	for ; i <= 3.25; i += 0.4 {
		massivA[count] = form(i)
		count += 1
	}
	return massivA
}
func zadB() []float64 {
	fmt.Println("ZAD B")
	var massivB = []float64{form(1.84), form(2.71), form(3.81), form(4.56), form(5.62)}
	return massivB
}
func form(x float64) float64 {
	return (math.Sqrt(math.Sqrt(math.Abs(x*x-2.5))) + math.Cbrt(math.Log10(x*x)))
}
