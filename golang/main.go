package main

import (
	"fmt"
	"math"
)

func main() {
	F4R(1.25, 3.25, 0.4)
	fmt.Println("---------------------")
	F4A([]float64{1.84, 2.71, 3.81, 4.56, 5.62})
}

func F4R(xn, xk, dx float64) {
	for x := xn; x <= xk; x += dx {
		var res float64 = math.Pow(math.Abs(math.Pow(x, 2)-2.5), 1.0/4.0) + math.Pow(math.Log(math.Pow(x, 2))/math.Log(math.E), 1.0/3.0)
		fmt.Println(res)
	}
}

func F4A(arr []float64) {
	for _, x := range arr {
		var res float64 = math.Pow(math.Abs(math.Pow(x, 2)-2.5), 1.0/4.0) + math.Pow(math.Log(math.Pow(x, 2))/math.Log(math.E), 1.0/3.0)
		fmt.Println(res)
	}
}
