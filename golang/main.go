package main

import (
	"fmt"
	"math"
)

func main() {
        fmt.Println("Num A")
	numA(0.4, 0.8, 3.2, 6.2, 0.6)
	fmt.Println("Mass")
	digit := [5]float64{4.48, 3.56, 2.78, 5.28, 3.21}
	Mass(0.4, 0.8, digit)
}
func lab(a float64, b float64, x float64) float64 {
	return (((math.Pow(a, x) - math.Pow(b, x)) / math.Log10(a/b)) * math.Pow((a*b), 1.0/3.0))
}
func numA(a float64, b float64, xn float64, xk float64, xd float64) {
	for x := xn; x < xk; x += xd {
		fmt.Println(lab(a, b, x))
	}
}
func Mass(a float64, b float64, digit [5]float64) {
	for i := 0; i < len(digit); i++ {
		fmt.Println(lab(a, b, digit[i]))
	}
}
