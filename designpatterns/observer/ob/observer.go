package ob

import (
	"context"
	"fmt"
)

type ObserverFace interface {
	Update(ctx context.Context, msg string)
}

type ObserverOne struct {
	Id string
}

func (o ObserverOne) Update(ctx context.Context, msg string) {
	fmt.Printf("Observer: %s, I get your msg: %s\n", o.Id, msg)
}

type ObserverTwo struct {
	Id string
}

func (o ObserverTwo) Update(ctx context.Context, msg string) {
	fmt.Printf("Observer: %s, I get your msg: %s\n", o.Id, msg)
}
