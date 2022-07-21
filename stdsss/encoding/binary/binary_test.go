package binary_test

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"testing"
)

func TestBinary(t *testing.T) {
	buf := new(bytes.Buffer)
	s := "Hello World"
	var version uint8 = 1
	var len uint32 = uint32(len(s))

	err := binary.Write(buf, binary.LittleEndian, &version)
	binary.Write(buf, binary.LittleEndian, &len)
	binary.Write(buf, binary.LittleEndian, []byte(s))

	// end := "我end"
	// binary.Write(buf, binary.LittleEndian, []byte(end))

	arr := []uint8{3, 4}
	binary.Write(buf, binary.LittleEndian, arr)
	arrn := binary.Size(arr)

	type user struct {
		Name string `json:"name"`
	}
	buser, _ := json.Marshal(user{Name: "Mike"})
	n := binary.Size(buser)
	fmt.Println("size: ")
	fmt.Println(n)
	binary.Write(buf, binary.LittleEndian, buser)

	if err != nil {
		fmt.Println("binary.Write failed:", err)
	}
	fmt.Printf("% x\n", buf.Bytes())

	// 类型尺寸
	head := make([]byte, 5)
	// 这个地方只是借用一下了buf, 实际上这可以是已经传输的conn对象
	io.ReadFull(buf, head)
	buf2 := bytes.NewBuffer(head)

	var v uint8
	var l uint32
	binary.Read(buf2, binary.LittleEndian, &v)
	binary.Read(buf2, binary.LittleEndian, &l)

	fmt.Println(v)
	fmt.Println(l)

	data := make([]byte, l)
	io.ReadFull(buf, data)

	fmt.Println(string(data))

	// ends := make([]byte, 3)
	// io.ReadFull(buf, ends)
	// fmt.Println(string(ends))

	a := make([]byte, arrn)
	io.ReadFull(buf, a)
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := make([]byte, n)
	io.ReadFull(buf, b)
	var uu user
	json.Unmarshal(b, &uu)

	fmt.Println(uu)
}
