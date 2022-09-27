package exercise_test

import (
	"fmt"
	"io"
	"os"
	"testing"
)

type WrapperWriter struct {
	writer  io.Writer
	counter int64
}

func (w *WrapperWriter) Write(p []byte) (n int, err error) {
	n, err = w.writer.Write(p)
	w.counter += int64(n)
	return n, err
}

// CountingWriter
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	newWriter := &WrapperWriter{writer: w}
	return newWriter, &newWriter.counter
}

func TestExercise72(t *testing.T) {
	writer, counter := CountingWriter(os.Stdout)
	fmt.Fprintf(writer, "Hello\n")
	fmt.Println(*counter)
	fmt.Fprintf(writer, "World\n")
	fmt.Println(*counter)
}
