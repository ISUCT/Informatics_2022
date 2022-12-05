package main

func aliveCellsCount(grid [][]string, rows, columns int) [][]int {
	newGrid := make([][]string, rows+1)
	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			newGrid[i][j] = grid[i-1][j-1]
		}
	}

	aliveCellsCount := make([][]int, rows)
	for i := 1; i <= rows; i++ {
		for j := 1; j <= columns; j++ {
			// Это лучшее, что я могла придумать в 11 вечера
			aliveCellsCount[i-1][j-1] =
				isCellAlive(newGrid[i+1][j]) + isCellAlive(newGrid[i][j+1]) + isCellAlive(newGrid[i+1][j+1]) +
					isCellAlive(newGrid[i-1][j]) + isCellAlive(newGrid[i][j-1]) + isCellAlive(newGrid[i-1][j-1]) +
					isCellAlive(newGrid[i+1][j-1]) + isCellAlive(newGrid[i-1][j+1])
		}
	}

	return aliveCellsCount
}

func isCellAlive(element string) int {
	if element == "#" {
		return 1
	}
	return 0
}

func main() {
	// n - width, m - height of a grid, s - steps
	var n, m, s int
	n, m, s = 5, 5, 4
	grid := make([][]string, m)
	for i := 0; i < n; i++ {
		grid[i] = make([]string, n)
	}

	grid = [][]string{{"....."}, {"..#.."}, {"#.#.."}, {".##.."}, {"....."}}

	for ; s > 0; s -= 1 {
		aliveCellsCount := aliveCellsCount(grid, m, n)
		for indexRow := 0; indexRow <= m; indexRow += 1 {
			for indexColumn := 0; indexColumn <= n; indexColumn += 1 {
				element := aliveCellsCount[indexRow][indexColumn]
				if grid[indexRow][indexColumn] == "." {
					if element == 3 {
						grid[indexRow][indexColumn] = "#"
					}
				} else if grid[indexRow][indexColumn] == "#" {
					if element != 2 && element != 3 {
						grid[indexRow][indexColumn] = "."
					}
				}
			}
		}
	}
}
