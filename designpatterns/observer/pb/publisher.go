package pb

import (
	"Gossssss/designpatterns/observer/ob"
	"context"
)

type Publisher interface {
	Subscribe(ctx context.Context, observer ob.ObserverFace)
	Notify(ctx context.Context, msg string)
	//  DeSubscribe(ctx context.Context, observer ObserverFace)
}

type SomeObject struct {
	Observers []ob.ObserverFace
	Name      string
	State     uint8
}

func NewObject(name string, state uint8) Publisher {
	return &SomeObject{
		Observers: make([]ob.ObserverFace, 0),
		Name:      name,
		State:     state,
	}
}

func (o *SomeObject) Subscribe(ctx context.Context, observer ob.ObserverFace) {
	o.Observers = append(o.Observers, observer)
}

func (o *SomeObject) Notify(ctx context.Context, msg string) {
	for _, observer := range o.Observers {
		observer.Update(ctx, msg)
	}
}
