package math_test

import (
	"fmt"
	"math"
	"strconv"
	"testing"
)

func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}

func Roundn(f float64, n int) float64 {
	floatStr := fmt.Sprintf("%."+strconv.Itoa(n)+"f", f)
	inst, _ := strconv.ParseFloat(floatStr, 64)
	return inst
}

func TestMath(t *testing.T) {
	var f float64 = 12.03
	fmt.Println(Round(f, 5))
	fmt.Println(Roundn(f, 5))
}
