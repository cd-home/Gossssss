package file

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestReadFileByBytes(t *testing.T) {

	file, err := os.Open("./URL.ini")

	if err != nil {
		permission := os.IsPermission(err)
		t.Log(permission)
		return
	}

	defer file.Close()
	// 按照字节读取
	for {
		buf := make([]byte, 11)
		n, err := file.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		// 	读取完
		if err == io.EOF || n == 0 {
			break
		}
		fmt.Println(string(buf[:n]))
	}

}

func TestReadFileByLine(t *testing.T) {
	file, err := os.Open("./URL.ini")

	if err != nil {
		permission := os.IsPermission(err)
		fmt.Println(permission)
		return
	}
	// 按照行读取
	reader := bufio.NewReader(file)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}

	// 改变文件偏移
	_, _ = file.Seek(0, io.SeekStart)

	scanner := bufio.NewScanner(file)
	// 定制分割符号
	//scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {})

	for scanner.Scan() {
		// 默认按照行读
		fmt.Println(scanner.Text())
	}
}
