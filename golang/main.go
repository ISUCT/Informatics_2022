package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("ZAD A")
	var i float64 = 1.25
	for i = 1.25; i <= 3.25; i += 0.4 {
		fmt.Println(form(i))
	}
	fmt.Println("ZAD B")
	fmt.Println(form(1.84))
	fmt.Println(form(2.71))
	fmt.Println(form(3.81))
	fmt.Println(form(4.56))
	fmt.Println(form(5.62))
}

func form(x float64) float64 {
	return (math.Sqrt(math.Sqrt(math.Abs(x*x-2.5))) + math.Cbrt(math.Log10(x*x)))
}
