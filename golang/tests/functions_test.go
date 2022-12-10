package internal_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"isuct.ru/informatics2022/internal"
)

type testCoordinates struct {
	row, column   int
	width, height int
}

func (s *testCoordinates) makeNewEntry(row, column, width, height int) {
	s.row = row
	s.column = column
	s.width = width
	s.height = height
}

// Обычный кейс, должно вернуть верное количество соседних клеток
func TestGetCoordinatesOfNeighborCellsNormal(t *testing.T) {
	testCase := testCoordinates{}
	testCase.makeNewEntry(5, 2, 10, 12)

	got := len(internal.GetCoordinatesOfNeighborCells(testCase.row, testCase.column, testCase.width, testCase.height))
	want := 8
	assert.Equal(t, want, got, fmt.Sprintf("got: %d, want: %d", got, want))
}

// Кейс, где соседних клеток нет
func TestGetCoordinatesOfNeighborCellsNoNeighbors(t *testing.T) {
	testCase := testCoordinates{}
	testCase.makeNewEntry(8, 7, 3, 6)

	got := len(internal.GetCoordinatesOfNeighborCells(testCase.row, testCase.column, testCase.width, testCase.height))
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

	testCase := testCoordinates{}
	testCase.makeNewEntry(1, 1, 3, 4)

	got := internal.CountAliveCells(
		gameOfLifeGrid,
		internal.GetCoordinatesOfNeighborCells(testCase.row, testCase.column, testCase.width, testCase.height),
	)
	want := 2
	assert.Equal(t, want, got, fmt.Sprintf("got: %d, want: %d", got, want))
}
