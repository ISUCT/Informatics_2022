package main

import (
	"fmt"
	"math"
)

func main() {
	taskA(2, 3, 0.11)
	taskB(2, 3)
}

func task(x, a, b float64) float64 {
	return (math.Asin(math.Pow(x, a)) + math.Acos(math.Pow(x, b)))
}

func taskA(a, b, x float64) {
	fmt.Println(`Part A`)
	for x <= 0.36 {
		fmt.Println(task(x, a, b))
		x += 0.05
	}
}

func taskB(a, b float64) {
	fmt.Println(`Part B`)
	fmt.Println(task(0.08, a, b))
	fmt.Println(task(0.26, a, b))
	fmt.Println(task(0.35, a, b))
	fmt.Println(task(0.41, a, b))
	fmt.Println(task(0.53, a, b))
}
