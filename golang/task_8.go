package main

import (
	"fmt"
	"sort"
)

func main() {
	answer := 0
	arr := [3][5]int{
		{1, 2, 3, 4, 5},
		{5, 6, 7, 8, 9},
		{20, 21, 34, 56, 100},
	}
	for _, n_1 := range arr {
		sort.Ints(n_1[:])
		answer += n_1[0]
	}
	fmt.Println(answer)
}
