package internal

import (
	"bufio"
	"os"
	"strings"
)

// Считывание слайса с клавиатуры
func ReadUserInput(width, height int, grid [][]string) {
	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < height; i++ {
		oneLineScan, err := reader.ReadString('\n')
		CheckForError(err)

		oneLineScan = strings.TrimSuffix(oneLineScan, "\n")
		inputtedSlice := strings.Split(oneLineScan, "") // Побуквенное разбиение строки на слайс

		for j := 0; j < width; j++ {
			grid[i][j] = inputtedSlice[j]
		}
	}
}
