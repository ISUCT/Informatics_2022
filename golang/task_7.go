package main

import "fmt"

func main() {
	fmt.Println(find_all([]int{6, 9, 3, 4, 3, 82, 11}, 3))
}

func find_all(arr []int, n int) (result []int) {
	for index, value := range arr {
		if value == n {
			result = append(result, index)
		}
	}
	return
}
