package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	flag.Parse()
	if len(flag.Args()) == 0 {
		return
	}
	for _, file := range flag.Args() {
		err := ReadeFileByline(file)
		if err != nil {
			fmt.Println(err)
		}

		err = ReadFileByWord(file)
		if err != nil {
			fmt.Println(err)
		}

		err = ReadFileByChar(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func ReadeFileByline(path string) error {
	var err error
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)

	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error read")
			break
		}
		fmt.Printf(line)
	}
	return nil
}

func ReadFileByWord(path string) error {
	var err error
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error")
			return err
		}
		r := regexp.MustCompile("[^\\s]+")
		words := r.FindAllString(line, -1)
		for i := 0; i < len(words); i++ {
			fmt.Println(words[i])
		}
	}
	return nil
}

func ReadFileByChar(path string) error {
	var err error
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()
	r := bufio.NewReader(f)
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error")
			return err
		}
		for _, x := range line {
			fmt.Println(string(x))
		}
	}
	return nil
}