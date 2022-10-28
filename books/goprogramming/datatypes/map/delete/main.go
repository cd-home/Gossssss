package main

import (
	"fmt"
)

func main() {
	m := make(map[string]any)
	m["name"] = "yao"
	m["age"] = 28
	m["sex"] = 0

	fmt.Println(m)

	// 获取key 个数
	fmt.Println(len(m))

	// 注意, 删除元素，map 不会释放这块的内存
	// 所在大规模的map缓存数据, 可能存在内存泄漏问题
	delete(m, "sex")
	// 删除操作是安全的， 可以删除不存在的key
	delete(m, "no exist key")

	fmt.Println(m)

	fmt.Println(len(m))
}
