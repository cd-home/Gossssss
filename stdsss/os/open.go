package main

import (
	"os"
	"testing"
)

func TestOpen(t *testing.T) {
	// 通常情况下，使用如下的文件操作方式， 注意默认是只读模式
	// 文件描述符
	file, err := os.Open("./URL.ini")

	if err != nil {
		return
	}

	defer file.Close()

	// 底层的文件描述符：一个非负整数（通常是小整数），来指代打开的文件
	//fmt.Println(file.Fd())
}

func TestOpenFile(t *testing.T) {

	// flag 打开文件模式 perm 文件权限
	openFile, err := os.OpenFile("os.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	// 	永远要记得关闭文件描述符， 通常这种错误不需要检测
	defer openFile.Close()

	// 	写入文件， 并不一定是立即刷新到磁盘
	_, _ = openFile.Write([]byte(`{"Name": "GodYao"}`))

	// 可以主动刷新， 或者设置打开文件的时候 os.O_SYNC
	_ = openFile.Sync()

	// 移除文件或者文件夹
	//_ = os.Remove("os.txt")

}
