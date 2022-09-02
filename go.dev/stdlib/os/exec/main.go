package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func ExecCommandCombined() {
	cmd := exec.Command("ls", "-l")
	// 执行命令, 并且返回标志输出和错误的合并
	output, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(output))
}

func ExecCommandEnviron() {
	cmd := exec.Command("pwd")
	// 返回环境变量
	envs := cmd.Environ()
	fmt.Println(envs)
}

func ExecCommandByStart() {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	err := cmd.Start()
	if err != nil {
		log.Println(err)
		return
	}
	err = cmd.Wait()
	if err != nil {
		log.Println(err)
		return
	}
}

func ExecCommandByRun() {
	cmd := exec.Command("ls")
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		log.Println(err)
		return
	}
}

func ExecCommandStdinPip() {
	cmd := exec.Command("grep", "values")
	// stdinPip 连接到 启动的命令的 标准输入
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		defer stdin.Close()
		io.WriteString(stdin, "values written to stdin \n are passed to cmd standard input")
	}()
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(output))
}

func ExecCommandStdoutPip() {
	// From
	cmdEcho := exec.Command("echo", "Hello World \n Hello You")
	stdout, err := cmdEcho.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmdEcho.Start(); err != nil {
		log.Fatal(err)
	}

	// To
	cmdGrep := exec.Command("grep", "You")
	stdin, err := cmdGrep.StdinPipe()
	if err != nil {
		log.Println(err)
		return
	}
	go func() {
		defer stdin.Close()
		io.Copy(stdin, stdout)
	}()
	data, err := cmdGrep.Output()
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(data))
}

func main() {
	//ExecCommandCombined()
	//ExecCommandEnviron()
	//ExecCommandByStart()
	//ExecCommandByRun()
	//ExecCommandStdinPip()
	ExecCommandStdoutPip()
}
