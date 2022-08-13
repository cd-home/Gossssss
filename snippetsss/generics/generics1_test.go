package generics_test

import (
	"testing"
)

type Number interface {
	~int64 | ~float64
}

func SumNumbers[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func TestGenericsNumber(t *testing.T) {
	m := map[string]int64{
		"a": 10,
		"b": 20,
	}
	t.Log(SumNumbers(m))

	m2 := map[string]float64{
		"d": 10.21,
		"e": 20.11,
	}
	t.Log(SumNumbers(m2))
}

func Add[T Number](a, b T) T {
	return a + b
}

func TestGenericsAdd(t *testing.T) {
	t.Log(Add[int64](1, 2))
	t.Log(Add[float64](1.2, 2.1))
}
