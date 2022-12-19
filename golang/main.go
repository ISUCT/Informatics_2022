package main

import (
	"fmt"

	"isuct.ru/informatics2022/internal"
)

func printSlice(grid [][]string, width, height int) {
	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			fmt.Print(grid[i][j])
		}
		fmt.Print("\n")
	}
}

func main() {
	// s - steps, n and m - width and height of a grid
	var s, n, m int
	_, err := fmt.Scanln(&s, &n, &m) // Для упрощения кода s, n и m вводятся на одной строке
	internal.CheckForError(err)

	grid := internal.InitializeGrid(n, m)
	internal.ReadUserInput(n, m, grid)

	secondGrid := internal.InitializeGrid(n, m)

	for ; s > 0; s -= 1 {
		internal.MakeASimulationStep(grid, secondGrid, n, m)
		internal.SwapGrids(&grid, &secondGrid)
	}

	printSlice(grid, n, m)
}
