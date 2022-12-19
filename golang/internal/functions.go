package internal

import (
	"log"
)

func CheckForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func InitializeGrid(width, height int) [][]string {
	grid := make([][]string, height)
	for i := 0; i < width; i++ {
		grid[i] = make([]string, width)
	}
	return grid
}

func GetCoordinatesOfNeighborCells(row, column, width, height int) (inBoundsCoordinates [][]int) {
	var allPossibleCoordinates = [][]int{
		{row + 1, column}, {row, column + 1}, {row + 1, column + 1},
		{row - 1, column}, {row, column - 1}, {row - 1, column - 1},
		{row + 1, column - 1}, {row - 1, column + 1},
	}
	// Проверка на выход за пределы слайса
	for _, coordinates := range allPossibleCoordinates {
		w, h := coordinates[0], coordinates[1]
		if (h < 0 || h > height-1) || (w < 0 || w > width-1) {
			continue
		} else {
			// Если координаты не указывают на элемент за пределами слайса, то добавляем их в inBoundsCoordinates
			inBoundsCoordinates = append(inBoundsCoordinates, coordinates)
		}
	}
	return inBoundsCoordinates
}

func CountAliveCells(grid [][]string, coordinatesOfNeighborCells [][]int) (alive int) {
	for _, coords := range coordinatesOfNeighborCells {
		i, j := coords[0], coords[1]
		if grid[i][j] == "#" {
			alive += 1
		}
	}
	return alive
}

func MakeASimulationStep(grid, secondGrid [][]string, width, height int) {
	for i := 0; i < height; i += 1 {
		for j := 0; j < width; j += 1 {
			alive := CountAliveCells(grid, GetCoordinatesOfNeighborCells(i, j, width, height))
			if (grid[i][j] == ".") && (alive == 3) {
				secondGrid[i][j] = "#"
			} else if (grid[i][j] == "#") && (alive == 2 || alive == 3) {
				secondGrid[i][j] = "#"
			} else {
				// Все символы, кроме #, будут интерпретироваться как точка
				secondGrid[i][j] = "."
			}
		}
	}
}

func SwapGrids(grid, secondGrid *[][]string) {
	*grid, *secondGrid = *secondGrid, *grid
}
