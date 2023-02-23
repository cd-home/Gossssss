package tests

import (
	"fmt"
	"testing"
)

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(2, -2)
	if ans != -2 {
		t.Errorf("IntMin(2, -2) = %d, but want -2", ans)
	}
}

func TestIntMinTableDriven(t *testing.T) {
	tests := []struct {
		a, b int
		want int
	}{
		{0, 1, 0},
		{1, 0, 0},
		{2, -2, -2},
	}

	for _, tt := range tests {
		testName := fmt.Sprintf("%d, %d", tt.a, tt.b)
		t.Run(testName, func(t *testing.T) {
			ans := IntMin(tt.a, tt.b)
			if ans != tt.want {
				t.Errorf("Got %d, But want %d", ans, tt.want)
			}
		})
	}
}
