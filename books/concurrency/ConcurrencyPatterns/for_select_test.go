package patterns_test

import (
	"context"
	"fmt"
	"testing"
	"time"
)

func TestForSelectSendDataToChan(t *testing.T) {
	// for-select
	/*
		for {
			select {}
		}*/
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()
	producer := func(ctx context.Context) <-chan string {
		stringStream := make(chan string)
		go func(ctx context.Context) {
			defer close(stringStream)
			// select send data to chan
			// and give timeout exit
			for _, s := range []string{"a", "b", "c", "d", "e"} {
				select {
				case <-ctx.Done():
					fmt.Println("timeout")
					return
				case stringStream <- s:
					time.Sleep(time.Second)
				}
			}
		}(ctx)
		return stringStream
	}
	consumer := func(data <-chan string) {
		for s := range data {
			fmt.Println(s)
		}
	}
	consumer(producer(ctx))
}

func TestForeverLoopSelect(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	producer := func() <-chan int {
		results := make(chan int)
		go func() {
			defer close(results)
			for i := 0; i < 10; i++ {
				results <- i
			}
		}()
		return results
	}
	consumer := func(ctx context.Context, data <-chan int) {
		for {
			// check. let select simple
			select {
			// but we give timeout exit
			case <-ctx.Done():
				fmt.Println("timeout")
				return
			default:

			}
			// task
			d, ok := <-data
			if ok {
				fmt.Println(d)
				time.Sleep(time.Second) // mock io time
			} else {
				fmt.Println("chan close")
			}
		}
	}
	old := time.Now()
	consumer(ctx, producer())
	duration := time.Since(old)
	fmt.Println(duration)
}

func TestForeverLoopSelectDefault(t *testing.T) {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*5)
	defer cancelFunc()
	producer := func() <-chan int {
		results := make(chan int)
		go func() {
			defer close(results)
			for i := 0; i < 10; i++ {
				results <- i
			}
		}()
		return results
	}
	consumer := func(ctx context.Context, data <-chan int) {
		for {
			select {
			// but we give timeout exit
			case <-ctx.Done():
				fmt.Println("timeout")
				return
			default:
				// 检测通道是否关闭
				d, ok := <-data
				if ok {
					fmt.Printf("%v, %d\n", ok, d)
					time.Sleep(time.Second) // mock io time
				} else {
					fmt.Println("chan close")
				}
			}
		}
	}
	consumer(ctx, producer())
}
