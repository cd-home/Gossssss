package main

import "fmt"

func main() {
	// 关联数据类型、或者称为映射、字典, 也是动态增长的
	m := make(map[string]int)

	// set 语法
	m["a"] = 1
	m["b"] = 2

	fmt.Println(m)

	// 返回key-value pair
	fmt.Println(len(m))

	// 获取key-value, 不存在也不会报错, 并且nil map是可读的
	v := m["a"]
	fmt.Println(v)

	// 删除key
	delete(m, "a")

	fmt.Println(m)

	// ok 模式 检测某个key 是否存在map 中
	if v, ok := m["a"]; ok {
		fmt.Println(v)
	} else {
		// 不存在返回value类型的零值
		fmt.Println("No a")
	}

	// 字面量, 声明并且初始化
	mps := map[string]string{"name": "yao", "age": "28"}
	fmt.Println(mps)
}
