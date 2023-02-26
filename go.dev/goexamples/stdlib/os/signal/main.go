package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	sigs := make(chan os.Signal, 1)
	done := make(chan struct{}, 1)

	// 监听信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Println(sig)
		done <- struct{}{}
	}()
	fmt.Println("awaiting signal...")
	<-done
	fmt.Println("exit...")
}
