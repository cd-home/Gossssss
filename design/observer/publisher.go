package observer

import "context"

type Publisher interface {
	Subscribe(ctx context.Context, observer ObserverFace)
	// DeSubscribe(ctx context.Context, observer ObserverFace)
	Notify(ctx context.Context, msg string)
}

type SomeObject struct {
	Observers []ObserverFace
	Name      string
	State     uint8
}

func NewObject(name string, state uint8) Publisher {
	return &SomeObject{
		Observers: make([]ObserverFace, 0),
		Name:      name,
		State:     state,
	}
}

func (o *SomeObject) Subscribe(ctx context.Context, observer ObserverFace) {
	o.Observers = append(o.Observers, observer)
}

func (o *SomeObject) Notify(ctx context.Context, msg string) {
	for _, observer := range o.Observers {
		observer.Update(ctx, msg)
	}
}
