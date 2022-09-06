package x

import (
	"fmt"
	"testing"
	"time"
)

func TestForRangeCopyValue(t *testing.T) {
	arr := []string{"a1", "b", "c", "d", "e", "f"}
	for i, v := range arr {
		// v is copy of arr[i]
		// &v 已经是最后一个迭代的值，所有看似不同的&v 都是指向最后一个
		t.Logf("%v: %v: %v\n", i, &v, &arr[i])
	}
}

func TestForRangeCopyValueFixe(t *testing.T) {
	arr := []string{"a1", "b", "c", "d", "e", "f"}
	for i, v := range arr {
		k := v
		t.Logf("%v: %v: %v: %v\n", i, &v, &k, &arr[i])
	}
}

func TestTemporaryPointer(t *testing.T) {
	store := make(map[int]*int)
	for i := 0; i < 5; i++ {
		store[i] = &i
	}
	// v is a value copy
	for _, v := range store {
		t.Log(*v)
	}
}

func TestForRangeChannel(t *testing.T) {
	ch := make(chan int, 3)
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		// close(ch)
	}()
	v, ok := <-ch
	t.Log(v, ok)
	v, ok = <-ch
	t.Log(v, ok)
	v, ok = <-ch
	t.Log(v, ok)
	v, ok = <-ch
	t.Log(v, ok)
	v, ok = <-ch
	t.Log(v, ok)
	v, ok = <-ch
	t.Log(v, ok)
}

func TestSelectChannel(t *testing.T) {
	ch := make(chan int)
	quit := make(chan bool)
	// 写
	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
			time.Sleep(time.Second * 2)
		}
		close(ch)    // 发送完成关闭channel
		quit <- true // 结束
	}()
	// 读
	for {
		// 监听channel数据
		select {
		case data := <-ch:
			fmt.Println(data)
		case <-quit:
			// 只能跳出select
			// break
			return
		default:
			fmt.Println("wait data")
			time.Sleep(time.Second)
		}
	}
}
