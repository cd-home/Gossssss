package unit

import "testing"

func add(x, y int) int {
	return x + y
}

func TestAdd(t *testing.T) {
	// table tests
	tests := []struct {
		name string
		want int
		x, y int
	}{
		{
			name: "one",
			want: 3,
			x:    1,
			y:    2,
		},
		{
			name: "two",
			want: 5,
			x:    2,
			y:    3,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if add(test.x, test.y) != test.want {
				t.Errorf("No Pass")
			}
		})
	}
}
