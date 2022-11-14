package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// 在Go中，函数运行失败时会返回错误信息，这些错误信息被认为是一种预期的值而非异常
	// 异常或者说错误 是预料不到的.
	// Go避免异常抛出的混乱, 而采用谁调用谁处理的原则

	resp, err := http.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println(err)
	}
	bs, _ := io.ReadAll(resp.Body)
	fmt.Println(string(bs))
	// 第一种错误，就是向上传播
	// 第二种错误，包装,自定义错误, 然后向上传播,
	// 第三种错误，打印错误，程序退出(一般用在main)
	// 第四种错误, 输出错误信息即可
	// 第五种错误，忽略错误 (尽量不要忽略错误)

	// 文件常见 io.EOF 固定"错误", 读取到文件最后
}
