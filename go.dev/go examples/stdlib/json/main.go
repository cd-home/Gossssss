package main

import (
	"encoding/json"
	"fmt"
	"time"
)
// 只有 可导出 的字段才会被 JSON 编码/解码
type User struct {
	Name string `json:"name"`
	// 如果是零值就不会输出
	Age      int       `json:"age,omitempty"`
	Sex      uint8     `json:"-"` // 忽略字段
	CreateAt time.Time `json:"create_at"`
	// 嵌套的结构体omitempty 如果不需要输出零值 需要使用指针
	Link *Link `json:"link,omitempty"`
}

type Link struct {
	Addr string `json:"addr"`
	Code int    `json:"code"`
}

func main() {
	u := User{
		Name:     "God Yao",
		Age:      12,
		Sex:      1,
		CreateAt: time.Now(),
	}
	fmt.Println(u)

	b, _ := json.Marshal(u)
	fmt.Println(string(b))

	byt := []byte(`{"name": "Jack", "age": 20}`)

	// 反序列化到struct
	u2 := User{}
	err := json.Unmarshal(byt, &u2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(u2)

	// 反序列化到map
	var m map[string]interface{}
	err = json.Unmarshal(byt, &m)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(m)

	fmt.Println(m["name"].(string))
	fmt.Println(m["age"].(float64))
}
