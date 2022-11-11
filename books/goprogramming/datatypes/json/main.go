package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type Movie struct {
	// 可导出, 才能序列化
	// 有默认的名字, 即是字段的名字, 亦可用json tag指定
	Title string
	// 结构体的成员Tag可以是任意的字符串面值，但是通常是一系列用空格分隔的key:"value"键值对序列
	// 最外层是反引号
	Year int `json:"released"`
	// omitempty 忽略零值
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	// JSON是对JavaScript中各种类型的值——字符串、数字、布尔值和对象
	mv := Movie{
		Title:  "明日",
		Year:   2022,
		Color:  false,
		Actors: []string{"a", "b", "c"},
	}
	bs, err := json.Marshal(mv)
	if err != nil {
		fmt.Println(err)
	}
	// json格式
	fmt.Printf("%s\n", bs)

	// 可以有缩进, 但只是用于展示
	bs, err = json.MarshalIndent(mv, "", "    ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", bs)

	// 反序列化
	// 先定义好结构体
	var mvv Movie
	err = json.Unmarshal(bs, &mvv)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(mvv)
}

// BigNumberDecode int64 精度丢失问题
func BigNumberDecode(str string) (map[string]interface{}, error) {
	var res map[string]interface{}

	d := json.NewDecoder(bytes.NewReader([]byte(str)))
	d.UseNumber() // 开启UseNumber属性即可，默认是关闭
	err := d.Decode(&res)
	if err != nil {
		return nil, err
	}

	return res, nil
}
