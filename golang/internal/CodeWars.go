package internal

import (
	"fmt"
	"strings"
)

func AllCodeWars() {
	// CodeWars 1
	fmt.Println(evenorodd(10))
	// CodeWars 2
	fmt.Println(countingsheep([]bool{true, true, true, true, true}))
	// CodeWars 3
	fmt.Println(countthemonkeys(10))
	// CodeWars 4
	fmt.Println(countingpaper(10, 10))
	//CodeWars 5
	fmt.Println(ishegonnasurvive(20, 10))
	//CodeWars 6
	fmt.Println(polishalphabet("Jędrzej Błądziński"))
	//CodeWars 7
	fmt.Println(findelement([]int{3, 3, 2, 2, 3}, 3))
	//CodeWars 8
	fmt.Println(sumofminimums([][]int{{1, 2, 3, 4, 5}, {5, 6, 7, 8, 9}, {20, 21, 34, 56, 100}}))
}

// CodeWars №1
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
func ishegonnasurvive(quantitybullet int, quantitydragon int) bool {
	fmt.Println("CodeWars №5 'Is he gonna survive?")
	if quantitybullet < quantitydragon*2 {
		return false
	} else {
		return true
	}
}

// CodeWars №6
func polishalphabet(s string) string {
	fmt.Println("CodeWars №6 Polish alphabet")
	var fstr string
	var array = strings.Split(s, "")
	var symbols = [][]string{{"ą", "a"}, {"ć", "c"}, {"ę", "e"}, {"ł", "l"}, {"ń", "n"}, {"ó", "o"}, {"ź", "z"}, {"ż", "z"}}
	for i := 0; i < len(array); i++ {
		var flag = true
		for j := 0; j < len(symbols); j++ {
			if strings.ToLower(array[i]) == symbols[j][0] {
				fstr = fstr + symbols[j][1]
				flag = false
			}
		}
		if flag {
			fstr = fstr + array[i]
		}
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
