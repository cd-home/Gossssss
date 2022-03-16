package main

import (
	"fmt"
	"os"
	"testing"
)

func TestStat(t *testing.T) {
	stat, err := os.Stat("URL.ini")
	if err != nil {
		return
	}
	dir := stat.IsDir()
	fmt.Println(dir)

	name := stat.Name()
	fmt.Println(name)

	mode := stat.Mode()
	fmt.Println(mode)

	size := stat.Size()
	fmt.Println(size)

	time := stat.ModTime()
	fmt.Println(time)

	// 底层数据来源
	sys := stat.Sys()
	fmt.Println(sys)
}
