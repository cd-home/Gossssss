package strings_test

import (
	"fmt"
	"strings"
	"testing"
)

func TestStringsbuilder(t *testing.T) {
	var buf strings.Builder
	buf.Grow(10)
	buf.WriteString("hello world")
	fmt.Println(buf.String())
}
