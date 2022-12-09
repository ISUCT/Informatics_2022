package main

import (
	"fmt"
	"math"
)

func main() {
	const a float64 = 2.0
	const b float64 = 1.1
	// Задание А
	task1(0.08, 0.2, 1.08, a, b)
	//Задание Б
	task2([]float64{0.1, 0.3, 0.4, 0.45, 0.65}, a, b)
	// CodeWars 1
	fmt.Println(evenorodd(10))
	// CodeWars 2
	fold := []bool{
		true, true, false}
	fmt.Println(countingsheep(fold))
	// CodeWars 3
	fmt.Println(countthemonkeys(10))
	// CodeWars 4
	fmt.Println(countingpaper(10, 10))
	//CodeWars 5
	fmt.Println(Ishegonnasurvive(20, 10))
	//CodeWars 6
}

func formula(x float64, a float64, b float64) float64 {
	var y float64 = (math.Log10(math.Abs(b*b-x*x)) / math.Log10(a)) / math.Pow(math.Abs(x*x-a*a), 0.2)
	return y
}

func task1(xn float64, xd float64, xk float64, a float64, b float64) {
	fmt.Println("Задание А")
	var x float64 = xn
	for ; x < xk; x = x + xd {
		fmt.Println(formula(x, a, b))
	}
}

func task2(Massiv []float64, a float64, b float64) {
	fmt.Println("Задание В")
	for i := 0; i < 5; i++ {
		Massiv[i] = formula(Massiv[i], a, b)
	}
	fmt.Println(Massiv)
}
func evenorodd(num int) string {
	fmt.Println(" CodeWars 'Even or Odd?'")
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func countingsheep(fold []bool) int {
	fmt.Println("CodeWars 'Counting sheep...'")
	var countsheep int = 0
	for _, value := range fold {
		if value {
			countsheep += 1
		}
	}
	return countsheep
}

func countthemonkeys(quantity int) []int {
	fmt.Println("CodeWars 'Count the Monkey!'")
	quantitymonkey := []int{}
	for i := 1; i < quantity+1; i++ {
		quantitymonkey = append(quantitymonkey, i)
	}
	return quantitymonkey
}

func countingpaper(n int, m int) int {
	fmt.Println("CodeWars 'Beginner Series #1 School Paperwork'")
	if n < 0 || m < 0 {
		return 0
	} else {
		return m * n
	}
}

func Ishegonnasurvive(quantitybullet int, quantitydragon int) bool {
	fmt.Println("CodeWars 'Is he gonna survive?")
	if quantitybullet < quantitydragon*2 {
		return false
	} else {
		return true
	}
}
