package interfaces_test

import (
	"bufio"
	"bytes"
	"fmt"
	"testing"
)

type WordCount int

// Write
func (wc *WordCount) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	// bufio.ScanLines
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*wc++
	}
	return len(p), nil
}

func TestExercise71(t *testing.T) {
	var wc WordCount
	wc.Write([]byte("Hello"))
	fmt.Println(wc)

	wc = 0
	name := "Dolly"
	fmt.Fprintf(&wc, "Hello %s", name)
	fmt.Println(wc)
}
