package main

import (
	"bufio"
	"fmt"
	"net/http"
)

func main() {
	resp, err := http.Get("https://www.baidu.com")

	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	//bs, _ := io.ReadAll(resp.Body)
	//fmt.Println(string(bs))

	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 5; i++ {
		fmt.Println(scanner.Text())
	}
}
