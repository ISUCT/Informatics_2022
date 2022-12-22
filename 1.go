package main

import (
	"fmt"
	"math"
)

func z1() {
	fmt.Printf("Задача 1: \n")
	for i := 0.26; i < 0.67; i += 0.08 {
		x := math.Pow(math.Asin(i), 2)
		y := math.Pow(math.Acos(i), 4)
		c := math.Pow((x + y), 3)
		fmt.Printf("%.2f \n", c)
	}
}

func z2() {
	fmt.Printf("Задача 2: \n")
	v := [5]float64{0.1, 0.35, 0.4, 0.55, 0.6}
	for i := 0; i < 5; i++ {
		x := math.Pow(math.Asin(v[i]), 2)
		y := math.Pow(math.Acos(v[i]), 4)
		c := math.Pow((x + y), 3)
		fmt.Printf("%.2f \n", c)
	}
}

func main() {
	defer z2()
	z1()
}
