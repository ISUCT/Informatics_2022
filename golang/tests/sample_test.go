package internal_test

import (
	"testing"

	"isuct.ru/informatics2022/internal"
)

func TestLog(t *testing.T) {
	result := internal.Log(100, 10)

	if result != 0.5 {
		t.Fatalf(`Summ(10,100) = %.4f, want 0.5, error`, result)
	}
}
