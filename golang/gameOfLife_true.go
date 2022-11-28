package main

import "fmt"

func gameOfLife(s int, width int, height int) {
	sample_grid := [][]bool{
		{false, false, false, false, false},
		{false, true, true, false, false},
		{false, true, true, false, false},
		{false, false, false, false, false},
		{false, false, false, false, false},
	}
	for i := 0; i < s; i++ {
		printGrid(width, height, gameOfLifeStep(width, height, sample_grid))
	}
}

func printGrid(width int, height int, grid [][]bool) {
	fmt.Println("Grid:")
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[x][y] {
				fmt.Printf("O")
			} else {
				fmt.Printf("#")
			}
		}
		fmt.Println()
	}
}

func gameOfLifeStep(width int, height int, grid [][]bool) [][]bool {
	var PosX int = 0
	var PosY int = 0
	var nb int
	for i := 0; i < width*height; i++ {
	PosX++
	if PosX > width {
		PosY++
		PosX = 0
	}
	if PosY > height {PosY = 0}
	if PosX > 0 && PosY > 0 && grid[PosX-1][PosY-1] {nb++}
	if PosX > 0 && grid[PosX-1][PosY] {nb++}
	if PosX > 0 && PosY < height && grid[PosX-1][PosY+1] {nb++}
	if PosY > 0 && grid[PosX][PosY-1] {nb++}
	if PosY < height-1 && grid[PosX][PosY+1] {nb++}
	if PosX < width-1 && PosY > 0 && grid[PosX+1][PosY-1] {nb++}
	if PosX < width-1 && grid[PosX+1][PosY] {nb++}
	if PosX < width-1 && PosY < height-1 && grid[PosX+1][PosY+1] {nb++}

	if nb == 3 {
		grid[PosX][PosY] = true
	}
	if !grid[PosX][PosY]{
		if nb == 3 {
			grid[PosX][PosY] = true
		}
	}}
	return grid
}
