package main

import (
	"log"
	"time"
)

func main() {
	t := time.Now()
	log.Println(t.Weekday(), t.Year(), t.Month(), t.Day())
	log.Println(t.Unix())

	//time.Sleep(time.Second * 2)
	//sub := t.Sub(time.Now())
	//log.Println(sub)

	log.Println(t.Local())
	log.Println(t.Format("2006/01/02 15:04:05"))

	loc , _ := time.LoadLocation("Asia/Shanghai")
	localTime := t.In(loc)
	log.Println(localTime)

	// 先解析字符串为时间
	d, _ := time.Parse(time.UnixDate, "Wed Feb 25 11:06:39 PST 2015")
	// 在格式化
	log.Println(d.Format("2006/01/02"))

	// 先解析字符串为时间
	d, _ = time.Parse("15:04", "13:36")
	// 在格式化
	log.Println(d.Hour(), d.Minute())

	// 先解析字符串为时间
	d, _ = time.Parse("02 January 2006", "12 January 2020")
	// 在格式化
	log.Println(d.Year(), d.Month(), d.Day())
	log.Println(d.Format("2006/01/02"))
}
