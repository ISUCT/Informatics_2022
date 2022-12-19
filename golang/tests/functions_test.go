package internal_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"isuct.ru/informatics2022/internal"
)

// Обычный кейс, должно вернуть верное количество соседних клеток
func TestGetCoordinatesOfNeighborCellsNormal(t *testing.T) {
	got := len(internal.GetCoordinatesOfNeighborCells(5, 2, 10, 12))
	want := 8
	assert.Equal(t, want, got, fmt.Sprintf("got: %d, want: %d", got, want))
}

// Кейс, где соседних клеток нет
func TestGetCoordinatesOfNeighborCellsNoNeighbors(t *testing.T) {
	got := len(internal.GetCoordinatesOfNeighborCells(8, 7, 3, 6))
	want := 0
	assert.Equal(t, want, got, fmt.Sprintf("got: %d, want: %d", got, want))
}

// Обычный кейс, должно вернуть верное количество живых клеток вокруг выбранной клетки
func TestCountAliveCellsNormal(t *testing.T) {
	var gameOfLifeGrid = [][]string{
		{".", ".", "."},
		{".", "#", "."},
		{"#", ".", "#"},
		{".", "#", "#"},
	}

	got := internal.CountAliveCells(gameOfLifeGrid, internal.GetCoordinatesOfNeighborCells(1, 1, 3, 4))
	want := 2
	assert.Equal(t, want, got, fmt.Sprintf("got: %d, want: %d", got, want))
}
