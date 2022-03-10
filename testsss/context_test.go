package testsss


import (
	"testing"
	"context"
	"time"
)

func DelayTask(ctx context.Context, stop chan struct{}) {
	go func(sign chan struct{}) {
		
	}(stop)
	select {
		case <-ctx.Done():
			stop <- struct{}{}
			return
	}
}


func TestContext(t *testing.T) {
	ctx, cancle := context.WithTimeout(context.Background(), time.Second * 3)
	defer cancle()
	select{
		case <-ctx.Done():
			t.Log("TimeOut")
		case <-time.After(time.Second):
			t.Log("TimeOut")
	}

}