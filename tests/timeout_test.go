package testsss

import (
	"context"
	"fmt"
	"testing"
	"time"
)

// TODO this
func Timeout(ctx context.Context) {
	// Warning chan cache must be >= 1
	ch := make(chan struct{}, 1)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- struct{}{}
	}()
	select {
		case <-ctx.Done():
			fmt.Println("Task is timed out")
			return
		case <-ch:
			fmt.Println("Task is successfully completed")
			return
		case <-time.After(time.Millisecond * 900):
			fmt.Println("Task is timed out")
			return
	}
}

func TestTimeOut(t *testing.T) {
	// Maybe this ctx pass by other
	ctx, cancle := context.WithTimeout(context.Background(), time.Second)
	defer cancle()
	Timeout(ctx)
}