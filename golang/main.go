package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	const a float64 = 4.1
	const b float64 = 2.7
	const xn float64 = 1.2
	const xk float64 = 5.2
	const xd float64 = 0.8
	// Задание А
	TaskA(a, b, xn, xk, xd)
	//Задание Б
	TaskB(a, b, []float64{1.9, 2.15, 2.34, 2.73, 3.16})
	// Задание 1
	fmt.Println(evenorodd(10))
	// Задание 2
	fmt.Println(countingsheep([]bool{true, true, true, true, true}))
	// Задание 3
	fmt.Println(countthemonkeys(10))
	// Задание 4
	fmt.Println(countingpaper(10, 10))
	// Задание 5
	fmt.Println(Ishegonnasurvive(20, 10))
	// Задание 6
	fmt.Println(polishalphabet("Jędrzej Błądziński"))
	// Задание 7
	fmt.Println(findelement([]int{3, 3, 2, 2, 3}, 3))
	// Задание 8
	fmt.Println(sumofminimums([][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 34, 56, 100}}))

}

// Формула для A и B

func Formula(x float64, a float64, b float64) float64 {
	var y float64 = (a*(math.Cbrt(x)) - (b * (math.Log(x) / math.Log(5)))) / (math.Log10(math.Abs(x - 1)))
	return y
}

// Задания A и B

func TaskA(a float64, b float64, xn float64, xk float64, xd float64) []float64 {
	var res []float64
	for i := xn; i < xk; i = i + xd {
		fmt.Println(Formula(i, a, b))
	}
	return res
}
func TaskB(a float64, b float64, x []float64) []float64 {
	var res []float64
	for i := 0; i < len(x); i++ {
		res = append(res, Formula(x[i], a, b))
	}
	return res
}

// Задание 1

func evenorodd(num int) string {
	fmt.Println(" Задание №1 'Even or Odd?'")
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

// Задание 2

func countingsheep(fold []bool) int {
	fmt.Println("Задание №2 'Counting sheep...'")
	var countsheep int = 0
	for _, value := range fold {
		if value {
			countsheep += 1
		}
	}
	return countsheep
}

// Задача 3

func countthemonkeys(quantity int) []int {
	fmt.Println("CodeWars №3 'Count the Monkey!'")
	quantitymonkey := []int{}
	for i := 1; i < quantity+1; i++ {
		quantitymonkey = append(quantitymonkey, i)
	}
	return quantitymonkey
}

// Задание 4

func countingpaper(n int, m int) int {
	fmt.Println("CodeWars №4 'Beginner Series #1 School Paperwork'")
	if n < 0 || m < 0 {
		return 0
	} else {
		return m * n
	}
}

// Задание 5

func Ishegonnasurvive(quantitybullet int, quantitydragon int) bool {
	fmt.Println("Задание №5 'Is he gonna survive?")
	if quantitybullet < quantitydragon*2 {
		return false
	} else {
		return true
	}
}

// Задание 6

func polishalphabet(s string) string {
	fmt.Println("Задание №6 Polish alphabet")
	var fstr string
	var flag bool = true
	var up = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	var lower = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var falsesigh = []string{"ą", "ć", "ę", "ł", "ń", "ó", "ś", "ź", "ż", " "}
	var truesigh = []string{"a", "c", "e", "l", "n", "o", "s", "z", "z", " "}
	var strmassiv = strings.Split(s, "")
	// Регистр
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

// Задание 7

func findelement(arraynum []int, num int) []int {
	fmt.Println("Задание №7 Find all occurrences of an element in an array")
	var numbers []int
	for i := 0; i < len(arraynum); i++ {
		if arraynum[i] == num {
			numbers = append(numbers, i)
		}
	}
	return numbers
}

// Задание 8

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
