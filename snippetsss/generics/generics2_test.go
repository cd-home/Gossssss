package generics_test

import "testing"

type E interface {
	~int64 | ~float64 | ~bool | ~string
}

type Object[T E] struct {
	Container []T
}

func (o *Object[T]) AppendObject(obj T) {
	o.Container = append(o.Container, obj)
}

func TestGenericsStruct(t *testing.T) {
	object := &Object[int64]{}
	object.AppendObject(1)
	object.AppendObject(2)
	object.AppendObject(3)
	t.Log(object)

	object2 := &Object[string]{}
	object2.AppendObject("a")
	object2.AppendObject("b")
	object2.AppendObject("c")

	t.Log(object2)

}
