package bufio_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"testing"
)

func TestBufio(t *testing.T) {
	// bytes
	bs := []byte("Hello\nWorld\nMy\nTime")

	// bytes reader [方便对bytes操作]
	bsreader := bytes.NewReader(bs)
	fmt.Println("size: ", bsreader.Size())
	bsreader.Seek(2, io.SeekStart)
	part := make([]byte, 2)
	bsreader.Read(part)

	fmt.Println(string(part))
	bsreader.Seek(0, io.SeekStart)

	fmt.Println("================================================")

	// bufio reader [方便对满足 io 接口的对象做操作, 适合文本io]
	bufioreader := bufio.NewReader(bsreader)
	data, _ := bufioreader.ReadBytes('\n')
	fmt.Println(string(data))

	fmt.Println("================================================")

	s, _ := bufioreader.ReadString('\n')
	fmt.Println(s)
}
