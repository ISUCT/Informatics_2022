package main

import (
	"testing"
)

type paperWorkTest struct {
	numberOfPeople int
	numberOfPages  int
	want           int
}

func TestPaperWork(t *testing.T) {
	var testCases = []paperWorkTest{
		{2, 5, 10},
		{0, 2, 0},
		{2, 0, 0},
		{-1, 5, 0},
		{1, -5, 0},
		{0, 0, 0},
	}

	for _, testCase := range testCases {
		got := PaperWork(testCase.numberOfPeople, testCase.numberOfPages)
		if got != testCase.want {
			t.Errorf("got %d, want %d", got, testCase.want)
		}
	}
}
