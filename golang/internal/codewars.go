package internal

import "math"

func EvenOrOdd(num int64) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}
func CountingSheeps(arrcs []bool) int {
	var csres int
	var i int
	for ; i < len(arrcs); i++ {
		if arrcs[i] {
			csres++
		}
	}
	return csres
}
func CountingMonkeys(n int) []int {
	var arrcm = make([]int, 0)
	for i := 1; i <= n; i++ {
		arrcm = append(arrcm, i)
	}
	return arrcm
}
func SchoolPaperwork(n int, m int) int {
	if (n >= 0) && (m >= 0) {
		return n * m
	} else {
		return 0
	}
}
func IsHe(bullets int, dragons int) string {
	if bullets >= dragons*2 {
		return "True"
	} else {
		return "False"
	}
}

//ПОЛЬСКАЯ КОРОВА НЕ ХОЧЕТ ТАНЦЕВАТЬ
/*func Polish(text string) []string {
	var res []string
	var vowels string = "ąćęłńóśźż"
	var standarts string = "acelnoszz"
	for i := 0; i < (len(text)); i++ {
		for k := 0; k < (len(vowels)); k++ {
			if text[i] == vowels[k] {
				res[i] = string(standarts[k])
			} else {
				res[i] = string(text[i])
			}
		}
	}
	return res
}
*/
func FindAll(arrFa []int, n int) []int {
	var res []int
	for i := 0; i < len(arrFa); i++ {
		if arrFa[i] == n {
			res = append(res, i)
		}
	}
	return res
}
func SumOfMin(arrA [][]int, m int, n int) int {
	var res int = 0
	for i := 0; i < (m); i++ {
		min := math.Inf(1)
		for k := 0; k < (n); k++ {
			if float64(arrA[i][k]) < min {
				min = float64(arrA[i][k])
			}
		}
		res = res + int(min)
	}
	return res
}

// text - string
// text[i] - byte
