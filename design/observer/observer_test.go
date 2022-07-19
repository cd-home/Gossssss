package observer

import (
	"context"
	"testing"
)

func TestObserver(t *testing.T) {
	object := NewObject("Goods", 1)

	object.Subscribe(context.Background(), &ObserverOne{Id: "10001"})
	object.Subscribe(context.Background(), &ObserverTwo{Id: "10002"})

	object.Notify(context.Background(), "Mac Pro 16 Publish!")
}
