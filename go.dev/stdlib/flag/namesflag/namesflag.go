package namesflag

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
	s.Names = append(s.Names, names...)
	return nil
}

func MoreFlags() {
	var manyNames NamesFlag
	minusk := flag.Int("k", 0, "An int")
	minusO := flag.String("o", "nike", "name")
	flag.Var(&manyNames, "names", "names list")

	flag.Parse()
	fmt.Println("-k", *minusk)
	fmt.Println("-o", *minusO)

	for idx, name := range manyNames.GetNames() {
		fmt.Println(idx, name)
	}

	for idx, args := range flag.Args() {
		fmt.Println(idx, args)
	}
}
