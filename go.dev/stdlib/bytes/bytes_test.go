package bytes_test

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBytes(t *testing.T) {
	// 新建一个buf 空间 来写入 bytes, 或者读取
	bs := []byte("Hello\nWorld\nMy\nTIme")
	bsbuf := bytes.NewBuffer(bs)

	bsbuf.Write([]byte("Hello\nWorld\n"))

	fmt.Println(bsbuf.String())

	// 初始化
	initbuf := new(bytes.Buffer)
	initbuf.WriteString("Hello World\n")
	initbuf.Write([]byte("GoGo"))
	fmt.Println(initbuf.String())
}
