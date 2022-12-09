package pool

import "context"

type Task interface {
	Run(ctx context.Context)
}

type RuntimeErr interface {
	Catch()
}
