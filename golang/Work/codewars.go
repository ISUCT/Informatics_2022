package main

func EvenOrOdd(number int) string {
	if number%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func CountingSheep(count []bool) int {
	k := 0
	for i := 0; i < len(count); i++ {
		if count[i] {
			k ++
		}
	}
	return k
}

func CountTheMonkeys(n int) []int {
	count := make([]int, 0, n)
	for i := 1; i <= n; i++ {
		count = append(count, i)
	}
	return count
}

func SchoolPaperwork(n, m int) int {
	if n < 0 || m < 0 {
		return 0
	} else {
		return n * m
	}
}

func IsHeGonnaSurvive(bullets, dragons int) string {
	if bullets/2 >= dragons {
		return "True"
	} else {
		return "False"
	}
}
