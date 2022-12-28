package internal

import (
	"fmt"
)

func GameOfLife(grid [][]byte, step int) {
	printGrid(grid, 0)
	for i := 1; i <= step; i++ {
		grid = printGrid(gol_Step(grid), i)
	}
}

func printGrid(grid [][]byte, step int) [][]byte {
	fmt.Println("Step ", step)
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			switch grid[x][y] {
			case 2:
				fmt.Print("# ")
			case 0:
				fmt.Print("+ ")
			case 1:
				fmt.Print("0 ")
			default:
				fmt.Print("? ")
			}
		}
		fmt.Println()
	}
	return grid
}

func checkNeighbors(grid [][]byte, PosX int, PosY int) int {
	var nb_list [][]int8 = [][]int8{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	nb := 0
	for nb_list_c := 0; nb_list_c < 8; nb_list_c++ {
		if grid[PosX+int(nb_list[nb_list_c][0])][PosY+int(nb_list[nb_list_c][1])] == 1 {
			nb++
		}
	}
	return nb
}

func gol_Step(grid [][]byte) [][]byte {
	overwrite_grid := make([][]byte, len(grid))
	for i := range grid {
		overwrite_grid[i] = make([]byte, len(grid[i]))
		copy(overwrite_grid[i], grid[i])
	}
	for PosX := 1; PosX < len(grid)-1; PosX++ {
		for PosY := 1; PosY < len(grid[0])-1; PosY++ {
			nb := checkNeighbors(grid, PosX, PosY)
			switch grid[PosX][PosY] {
			case 0:
				if nb == 3 {
					overwrite_grid[PosX][PosY] = 1
				} else {
					overwrite_grid[PosX][PosY] = 0
				}
			case 1:
				if nb == 2 || nb == 3 {
					overwrite_grid[PosX][PosY] = 1
				} else {
					overwrite_grid[PosX][PosY] = 0
				}
			}
		}
	}
	return overwrite_grid
}

func IsGridCorrect(grid [][]byte) bool {
	for i := 0; i < len(grid[0])-1; i++ {
		if grid[0][i] != byte(2) && grid[len(grid)-1][i] != byte(2) {
			return false
		}
	}
	for t := 1; t < len(grid)-2; t++ {
		if grid[t][0] != byte(2) && grid[t][len(grid[0])-1] != byte(2) {
			return false
		}
	}
	return true
}
