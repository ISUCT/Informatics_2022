package main

import (
	"fmt"
	"math"

	"isuct.ru/informatics2022/internal"
)

func main() {
	zada(2.2, 1.2, 0.48)
	zadb(1.2, 0.48, []float64{0.25, 0.36, 0.56, 0.94, 1.28}, 0)
	fmt.Println(internal.EvenOrOdd(5))
	internal.Countingsheep([]bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true})
	internal.CounttheMonkeys(4)
	fmt.Println(internal.BeginnerSeries(0, -1))
	fmt.Println(internal.Ishegonnasurvive(2, 4))

}

func zada(xk float64, a float64, b float64) {
	fmt.Println("Задание А")

	for x := 0.7; x < xk; x = x + 0.3 {
		var y float64 = (((math.Pow(b, 3)) + math.Pow(math.Sin(a*x), 2)) / ((math.Acos(x * b * x)) + (math.Pow(2.7183, -x/2))))
		fmt.Println(y)
	}

}
func zadb(a float64, b float64, numbersx []float64, c int64) {
	fmt.Println("Задание Б")

	for ; c <= 4; c = c + 1 {
		var y float64 = (((math.Pow(b, 3)) + math.Pow(math.Sin(a*numbersx[c]), 2)) / ((math.Acos(numbersx[c] * b * numbersx[c])) + (math.Pow(2.7183, -numbersx[c]/2))))
		fmt.Println(y)

	}
}
