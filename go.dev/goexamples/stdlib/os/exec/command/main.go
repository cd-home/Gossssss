package main

import (
	"fmt"
	"io"
	"os/exec"
)

func ExecCommand() {
	// Will LookPath
	// 通常执行环境变量中的可执行文件
	dateCmd := exec.Command("date")

	// 执行后需要显式输出, Output实际上就是指定 output 和执行 run
	output, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	fmt.Println(string(output))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

func PipLineExec() {
	grepCmd := exec.Command("grep", "hello")
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	bytes, _ := io.ReadAll(grepOut)
	grepCmd.Wait()
	fmt.Println("> grep hello")
	fmt.Println(string(bytes))
}

func BashExec() {
	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("->  ls -a -l -h")
	fmt.Println(string(lsOut))
}

func main() {
	ExecCommand()
}
