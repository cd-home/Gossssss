package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func ReadFileByLine(path string) error {
	var err error
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	r := bufio.NewReader(f)
	for {
		// 按照行读
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println(line)
	}
	return nil
}

func ReadFileByWord(path string) error {
	var err error
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Println(err)
		}
	}(f)
	words := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		words++
		fmt.Println(scanner.Text())
	}
	fmt.Println(words)
	return nil
}

func ReadFileByCustomDelim() error {
	buf := bytes.NewBufferString("1, 2, 3, 4")
	scanner := bufio.NewScanner(buf)
	onComma := func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		for i := 0; i < len(data); i++ {
			if data[i] == ',' {
				return i + 1, data[:i], nil
			}
		}
		if !atEOF {
			return 0, nil, nil
		}
		return 0, data, bufio.ErrFinalToken
	}
	scanner.Split(onComma)
	for scanner.Scan() {
		fmt.Println(strings.Trim(scanner.Text(), " "))
	}
	return nil
}

// Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer
// object, creating another object (Reader or Writer) that also implements
// the interface but provides buffering and some help for textual I/O.
// [文本处理的IO]
func main() {
	err := ReadFileByLine("./main.go")
	if err != nil {
		log.Println(err)
		return
	}
	err = ReadFileByWord("./main.go")
	if err != nil {
		log.Println(err)
		return
	}
	err = ReadFileByCustomDelim()
	if err != nil {
		return
	}
}
