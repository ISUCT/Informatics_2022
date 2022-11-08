package internal_test

import (
	"testing"

	"isuct.ru/informatics2022/internal"
)

func TestSumm(t *testing.T) {
	summ := internal.Summ(2, 3)
	if summ != 5 {
		t.Fatalf(`Summ(2,3) = %d, want 5, error`, summ)
	}
}
