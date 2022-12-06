package main

func sheeps(slice []bool) int {
	nmofsheeps := 0
	for i := 0; i < len(slice); i++ {
		if slice[i] {
			nmofsheeps++
		}
	}
	return nmofsheeps
}
