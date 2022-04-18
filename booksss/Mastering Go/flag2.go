package main

import (
	"flag"
	"fmt"
	"strings"
)

type NamesFlag struct {
	Names []string
}

func (s *NamesFlag) GetNames() []string {
	return s.Names
}

func (s *NamesFlag) String() string {
	return fmt.Sprint(s.Names)
}

func (s *NamesFlag) Set(v string) error {
	if len(s.Names) > 0 {
		return fmt.Errorf("no")
	}
	names := strings.Split(v, ",")
	for _, item := range names {
		s.Names = append(s.Names, item)
	}
	return nil
}


func main() {
	var manyNames  NamesFlag
	minusk := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "nike", "name")
	flag.Var(&manyNames, "names", "names list")

	flag.Parse()
	fmt.Println("-k", *minusk)
	fmt.Println("-o", *minusO)

	for i, item := range manyNames.GetNames() {
		fmt.Println(i ,item)
	}

	for index, val := range flag.Args() {
		fmt.Println(index, val)
	}

}
