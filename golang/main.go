package main

import (
	"fmt"
	"math"

	"isuct.ru/informatics2022/internal"
)

func main() {
	ZadanieA(1.35, 0.98, 4.24, 1.14)
	ZadanieB(1.35, 0.98, 0)
	internal.EvenOrOdd(4)
	internal.CountingSheep([]bool{true, true, true, false,
		true, true, true, true,
		true, false, true, false,
		true, false, false, true,
		true, true, true, true,
		false, false, true, true})
	internal.MonkeysCount(4)
	fmt.Println(internal.PaperworkCount(0, -1))
	fmt.Println(internal.HeroWithGunShootsDragons(2, 6))
	fmt.Println(internal.PolishCow("Grzegorz Brzęczyszczykiewicz"))
	fmt.Println(internal.FindAll([]int{1, 4, 5, 7, 2, 1, 4, 8, 8}, 4))
	fmt.Println(internal.SumOfMin([][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 34, 56, 100}}, 3, 5))

}

func ZadanieA(a float64, b float64, xk float64, x float64) {
	fmt.Println("Задание А")
	for ; x < xk; x += 0.62 {
		var y float64 = (math.Pow(a*x+b, 1.0/3)) / (math.Pow((math.Log10(x)), 2))
		fmt.Println(y)
	}
}
func ZadanieB(a float64, b float64, c int64) {
	fmt.Println("Задание Б")
	var numbersx []float64 = []float64{0.35, 1.28, 3.51, 5.21, 4.16}
	for ; c <= 4; c += 1 {
		var y float64 = (math.Pow(a*numbersx[c]+b, 1.0/3)) / (math.Pow((math.Log10(numbersx[c])), 2))
		fmt.Println(y)
	}
}
