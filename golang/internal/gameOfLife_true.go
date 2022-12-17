package internal

import "fmt"

func GameOfLife(s int, width int, height int, grid [][]byte) {
	for i := 0; i < s; i++ {
		grid = GameOfLifeStep(width, height, grid)
		printGrid(width, height, grid, i+1)
	}
}

func printGrid(width int, height int, grid [][]byte, step int) [][]byte {
	fmt.Println("Step ", step)
	for y := 0; y <= height+1; y++ {
		for x := 0; x <= width+1; x++ {
			switch grid[x][y] {
			case byte(1):
				fmt.Printf("O")
			case byte(0):
				fmt.Printf("#")
			case byte(2):
				fmt.Printf("+")
			}
		}
		fmt.Println()
	}
	return grid
}

func GameOfLifeStep(width int, height int, grid [][]byte) [][]byte {
	nb_grid := [][]int{
		{-1, -1},
		{-1, 0},
		{-1, 1},
		{0, -1},
		{0, 1},
		{1, -1},
		{1, 0},
		{1, 1},
	}
	var PosX int = 0
	var PosY int = 1
	var nb int
	var overwrite_pos [][]int
	for i := 0; i < width*height; i++ {
		PosX++
		if PosX > width {
			PosY++
			if PosY > height {
				PosY = 0
			}
			PosX = 1
		}
		for ia := 0; ia < len(nb_grid); ia++ {
			if grid[PosX+nb_grid[ia][0]][PosY+nb_grid[ia][1]] == 1 {
				nb++
			}
		}
		if grid[PosX][PosY] == 0 && nb == 3 {
			overwrite_pos = append(overwrite_pos, []int{PosX, PosY, 1})
		} else {
			if grid[PosX][PosY] == 1 && (nb == 3 || nb == 2) {
				overwrite_pos = append(overwrite_pos, []int{PosX, PosY, 1})
			} else {
				overwrite_pos = append(overwrite_pos, []int{PosX, PosY, 0})
			}
		}
		nb = 0
	}

	for ib := 0; ib < len(overwrite_pos); ib++ {
		grid[overwrite_pos[ib][0]][overwrite_pos[ib][1]] = byte(overwrite_pos[ib][2])
	}
	return grid
}
