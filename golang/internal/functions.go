package internal

import (
	"log"
)

func CheckForError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

// Возвращает координаты соседних клеток
func GetCoordinatesOfNeighborCells(row, column, width, height int) (coords [][]int) {
	var coordinatesOfNeighborCells = [][]int{
		{row + 1, column}, {row, column + 1}, {row + 1, column + 1},
		{row - 1, column}, {row, column - 1}, {row - 1, column - 1},
		{row + 1, column - 1}, {row - 1, column + 1},
	}
	// Проверка на выход за пределы слайса
	for _, coordinates := range coordinatesOfNeighborCells {
		w, h := coordinates[0], coordinates[1]
		if (h < 0 || h > height-1) || (w < 0 || w > width-1) {
			continue
		} else {
			coords = append(coords, coordinates)
		}
	}
	return coords
}

// Возвращает количество живых клеток среди соседних
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
				secondGrid[i][j] = "."
			}
		}
	}
}

func SwapGrids(grid, secondGrid *[][]string) {
	*grid, *secondGrid = *secondGrid, *grid
}
