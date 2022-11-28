package main

func evenOrOdd(num int) string {
	if num%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func countingSheep(arr []bool) int {
	var res int
	var i int
	for ; i < len(arr); i++ {
		if arr[i] && (arr[i] || !arr[i]) {
			res++
		}
	}
	return res
}

func monkeysCount(num int) []int {
	var res []int
	for n := 1; n <= num; n++ {
		res = append(res, n)
	}
	return res
}

func paperworkCount(n int, m int) int {
	if n < 0 || m < 0 {
		return 0
	}
	return n * m
}

func heroWithGunShootsDragons(ammo int, dragons int) bool {
	return ammo/2 >= dragons
}
