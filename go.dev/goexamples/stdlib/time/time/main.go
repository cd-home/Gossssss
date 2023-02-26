package main

import (
	"fmt"
	"time"
)

func main() {
	// date
	now := time.Now()
	then := time.Date(2019, 1, 12, 10, 20, 50, 0, time.UTC)
	diff := now.Sub(then)
	fmt.Println(diff)
	fmt.Println(then.Add(-diff))
	fmt.Println(then.Add(diff))

	// unix
	fmt.Println(time.Now())
	fmt.Println(time.Now().Unix())
	fmt.Println(time.Now().UnixNano())

	second := time.Now().Unix()
	fmt.Println(time.Unix(second, 0))

	// format parse
	fmt.Println(time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	fmt.Println(time.Now().Add(time.Hour))

	fmt.Println(time.Now().AddDate(0, 0, 1))

	duration, err := time.ParseDuration("1h30m")
	if err != nil {
		return
	}
	fmt.Println(duration)
}
