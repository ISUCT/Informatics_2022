package internal

import "fmt"

func EvenOrOdd(n int) string {
	if n%2 == 0 {
		return "even"
	} else {
		return "odd"
	}
}

func Countingsheep(Ovces []bool) {
	var c int64 = 0
	for i := 0; i < len(Ovces); i++ {
		if Ovces[i] == true {
			c++
		}
	}
	fmt.Println("Количество овец ", c)
}

func CounttheMonkeys(n int) {
	var obezyans = []int{}
	for i := 1; i <= n; i++ {
		obezyans = append(obezyans, int(i))
	}
	fmt.Println(obezyans)
}

func BeginnerSeries(n int, m int) int {
	if m < 0 || n < 0 {
		return 0
	} else {
		return m * n
	}
}

func Ishegonnasurvive(d int, b int) string {
	if (d*2)-b <= 0 {
		return "Hero Win"
	} else {
		return "Hero Died"
	}

}
