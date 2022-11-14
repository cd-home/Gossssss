package main

import "os"

func main() {

}

func DoFile(filenames []string) error {
	// 循环中调用defer需要注意
	for _, filename := range filenames {
		if err := doFile(filename); err != nil {
			return err
		}
	}
	return nil
}

func doFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	// logic
	return nil
}