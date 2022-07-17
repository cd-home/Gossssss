package fmt_test

import (
	"fmt"
	"testing"
)

func TestPaddingZero(t *testing.T) {
	i := 12.03
	fmt.Printf("%.4f\n", i)
}
