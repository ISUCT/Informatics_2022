package internal

import (
	"fmt"
)

func printGrid(grid [][]byte, step int) [][]byte {
	fmt.Println("Step ", step)
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			switch grid[y][x] {
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

func GameOfLife(grid [][]byte, step int) {
	if !(isGridCorrect(grid)) {
		fmt.Println("Grid check failed")
	}
	for i := 1; i <= step; i++ {
		printGrid(gol_Step(grid), i)
	}
}

func gol_Step(grid [][]byte) [][]byte {
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
	var overwrite_pos [][]int
	for PosY := 1; PosY < len(grid)-1; PosY++ {
		for PosX := 1; PosX < len(grid[0])-1; PosX++ {
			nb := 0
			for nb_list_c := 0; nb_list_c < 8; nb_list_c++ {
				if grid[PosX+int(nb_list[nb_list_c][0])][PosY+int(nb_list[nb_list_c][1])] == 1 {
					nb++
				}
			}
			switch grid[PosX][PosY] {
			case 0:
				if nb == 3 {
					overwrite_pos = append(overwrite_pos, []int{PosX, PosY, 1})
				} else {
					overwrite_pos = append(overwrite_pos, []int{PosX, PosY, 0})
				}
			case 1:
				if nb == 2 || nb == 3 {
					overwrite_pos = append(overwrite_pos, []int{PosX, PosY, 1})
				} else {
					overwrite_pos = append(overwrite_pos, []int{PosX, PosY, 0})
				}
			}
		}
	}
	for ow_c := 0; ow_c < len(overwrite_pos); ow_c++ {
		grid[overwrite_pos[ow_c][0]][overwrite_pos[ow_c][1]] = byte(overwrite_pos[ow_c][2])
	}
	return grid
}

func isGridCorrect(grid [][]byte) bool {
	for i := 0; i < len(grid[0])-1; i++ {
		if grid[0][i] != byte(3) && grid[len(grid)-1][i] != byte(3) {
			return false
		}
	}
	for t := 1; t < len(grid[1])-2; t++ {
		if grid[t][0] != byte(3) && grid[t][len(grid[0])-1] != byte(3) && len(grid[0]) == len(grid[t]) {
			return false
		}
	}
	return true
}
