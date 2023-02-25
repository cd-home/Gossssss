package main

import (
	"Gossssss/designpatterns/observer/ob"
	"Gossssss/designpatterns/observer/pb"
	"context"
)

func main() {
	object := pb.NewObject("Goods", 1)

	object.Subscribe(context.Background(), &ob.ObserverOne{Id: "10001"})
	object.Subscribe(context.Background(), &ob.ObserverTwo{Id: "10002"})

	object.Notify(context.Background(), "Mac Pro 16 Publish!")
}
