package main

import (
	"fmt"
	"math"
	"strings"
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
	fmt.Println(countingsheep([]bool{true, true, true, true, true}))
	// CodeWars 3
	fmt.Println(countthemonkeys(10))
	// CodeWars 4
	fmt.Println(countingpaper(10, 10))
	//CodeWars 5
	fmt.Println(Ishegonnasurvive(20, 10))
	//CodeWars 6
	fmt.Println(polishalphabet("Jędrzej Błądziński"))
	//CodeWars 7
	fmt.Println(findelement([]int{3, 3, 2, 2, 3}, 3))
	//CodeWars 8
	fmt.Println(sumofminimums([][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 34, 56, 100}}))
}

// Формула для Заданий А и Б
func formula(x float64, a float64, b float64) float64 {
	var y float64 = (math.Log10(math.Abs(b*b-x*x)) / math.Log10(a)) / math.Pow(math.Abs(x*x-a*a), 0.2)
	return y
}

// Задание А
func task1(xn float64, xd float64, xk float64, a float64, b float64) {
	fmt.Println("Задание А")
	var x float64 = xn
	for ; x < xk; x = x + xd {
		fmt.Println(formula(x, a, b))
	}
}

// Задание Б
func task2(Massiv []float64, a float64, b float64) {
	fmt.Println("Задание В")
	for i := 0; i < 5; i++ {
		Massiv[i] = formula(Massiv[i], a, b)
	}
	fmt.Println(Massiv)

	//CodeWars №1
}
func evenorodd(num int) string {
	fmt.Println(" CodeWars №1 'Even or Odd?'")
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// CodeWars №2
func countingsheep(fold []bool) int {
	fmt.Println("CodeWars №2 'Counting sheep...'")
	var countsheep int = 0
	for _, value := range fold {
		if value {
			countsheep += 1
		}
	}
	return countsheep
}

// CodeWars №3
func countthemonkeys(quantity int) []int {
	fmt.Println("CodeWars №3 'Count the Monkey!'")
	quantitymonkey := []int{}
	for i := 1; i < quantity+1; i++ {
		quantitymonkey = append(quantitymonkey, i)
	}
	return quantitymonkey
}

// CodeWars №4
func countingpaper(n int, m int) int {
	fmt.Println("CodeWars №4 'Beginner Series #1 School Paperwork'")
	if n < 0 || m < 0 {
		return 0
	} else {
		return m * n
	}
}

// CodeWars №5
func Ishegonnasurvive(quantitybullet int, quantitydragon int) bool {
	fmt.Println("CodeWars №5 'Is he gonna survive?")
	if quantitybullet < quantitydragon*2 {
		return false
	} else {
		return true
	}
}

// CodeWars №6
// challenge completed
func polishalphabet(s string) string {
	fmt.Println("CodeWars №6 Polish alphabet")
	var fstr string
	var flag bool = true
	var up = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var lower = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var falsesigh = []string{"ą", "ć", "ę", "ł", "ń", "ó", "ś", "ź", "ż", " "}
	var truesigh = []string{"a", "c", "e", "l", "n", "o", "s", "z", "z", " "}
	var strmassiv = strings.Split(s, "")
	//Проверка регистра
	for d := 0; d < len(strmassiv); d++ {
		for uplower := 0; uplower < len(up); uplower++ {
			if strmassiv[d] == up[uplower] {
				strmassiv[d] = lower[uplower]
			}
		}
	}
	//Исправление символов
	for i := 0; i < len(strmassiv); i++ {
		for eng := 0; eng < len(lower); eng++ {
			if strmassiv[i] == lower[eng] {
				flag = false
			}
		}
		if flag {
			for sigh := 0; sigh < len(falsesigh); sigh++ {
				if strmassiv[i] == falsesigh[sigh] {
					fstr = fstr + truesigh[sigh]
				}
			}
		} else {
			fstr = fstr + strmassiv[i]
		}
		flag = true
	}
	return fstr
}

// CodeWars №7
func findelement(arraynum []int, num int) []int {
	fmt.Println("CodeWars №7 Find all occurrences of an element in an array")
	var numbers []int
	for i := 0; i < len(arraynum); i++ {
		if arraynum[i] == num {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

// CodeWars №8
func sumofminimums(numarray [][]int) int {
	fmt.Println("CodeWars №8 'Sum of Minimums!'")
	var sumnum int = 0
	for i := 0; i < len(numarray); i++ {
		var minnum int
		minnum = numarray[i][0]
		for j := 1; j < len(numarray[i]); j++ {
			if numarray[i][j] < minnum {
				minnum = numarray[i][j]
			}
		}
		sumnum = sumnum + minnum
	}

	return sumnum
}
