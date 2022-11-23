package main

func evenOrOdd(x int) string {
	if x%2 == 0 {
		return "Even"
	} else {
		return "Odd"
	}
}

func countingSheep(x []bool) int {
	var r int
	var i int
	for ; i < len(x); i++ {
		if x[i] && (x[i] || !x[i]) {
			r++
		}
	}
	return r
}

func monkeysCount(num int) []int {
	var r []int
	for n := 1; n <= num; n++ {
		r = append(r, n)
	}
	return r
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
