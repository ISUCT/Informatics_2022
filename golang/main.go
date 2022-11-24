package main

import (
	"fmt"
	"math"
)

func main() {
	for x := 0.7; x <= 2.2; x += 0.3 {
		y := par(1.2, 0.48, 2.72, x)
		fmt.Println(y)
	}
	var users = [5]float64{0.25, 0.36, 0.56, 0.94, 1.28}
	for i := 0; i < len(users); i++ {
		x := users[i]
		y := par2(1.2, 0.48, 2.72, x)
		fmt.Println(y)
	}
}

func par(a, b, e, x float64) float64 {
	y := (math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos(x*b*x) + (math.Pow(e, -x/2)))
	return y
}
func par2(a, b, e, x float64) float64 {
	var y float64 = (math.Pow(b, 3) + math.Pow(math.Sin(a*x), 2)) / (math.Acos(x*b*x) + (math.Pow(e, -x/2)))
	return y
}
