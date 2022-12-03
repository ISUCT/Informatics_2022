package main

import (
	"fmt"
	"math"
)

func main() {
	const a float64 = 2.0
	const b float64 = 1.1
	// Задание А
	task1(a, b)
	//Задание Б
	fmt.Println(task2(a, b))
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
}

func formula(x float64, a float64, b float64) float64 {
	var y float64 = (math.Log10(math.Abs(b*b-x*x)) / math.Log10(a)) / math.Pow(math.Abs(x*x-a*a), 0.2)
	return y
}

func task1(a float64, b float64) {
	fmt.Println("Задание А")
	var x float64 = 0.08
	for ; x < 1.08; x = x + 0.2 {
		fmt.Println(formula(x, a, b))
	}
}

func task2(a float64, b float64) [5]float64 {
	fmt.Println("Задание В")
	var Massiv = [5]float64{formula(0.1, a, b), formula(0.3, a, b), formula(0.4, a, b), formula(0.45, a, b), formula(0.65, a, b)}
	return Massiv
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
