package main

func sumofmin(inpsl [][]int) int {
	sumofm := 0
	for i := 0; i < len(inpsl); i++ {
		minn := inpsl[i][0]
		for j := 0; j < len(inpsl[i]); j++ {
			if inpsl[i][j] < minn {
				minn = inpsl[i][j]
			}
		}
		sumofm = sumofm + minn
	}
	return sumofm
}
