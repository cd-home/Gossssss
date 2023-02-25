package design

import (
	"testing"
)

// Operator 抽象公共接口
type Operator interface {
	SetA(a int)
	SetB(b int)
	Result() int
}

// OperatorFactoryFace 抽象工厂
type OperatorFactoryFace interface {
	Create() Operator
}

// OperatorBase 公共类，封装公共方法
type OperatorBase struct {
	A int
	B int
}

func (o *OperatorBase) SetA(a int) {
	o.A = a
}

func (o *OperatorBase) SetB(b int) {
	o.B = b
}

type PlusOperatorFactory struct{}
type PlusOperator struct {
	*OperatorBase
}
// Result 实现自己的Result逻辑
func (o PlusOperator) Result() int {
	return o.A + o.B
}
func (PlusOperatorFactory) Create() Operator {
	return &PlusOperator{OperatorBase: &OperatorBase{}}
}

type MinusOperatorFactory struct{}
type MinusOperator struct {
	*OperatorBase
}
// Result 实现自己的Result逻辑
func (o MinusOperator) Result() int {
	return o.A / o.B
}
func (MinusOperatorFactory) Create() Operator {
	return &MinusOperator{OperatorBase: &OperatorBase{}}
}

func compute(factory OperatorFactoryFace, a, b int) int {
	op := factory.Create()
	op.SetA(a)
	op.SetB(b)
	return op.Result()
}

func TestFactoryMethod(t *testing.T) {
	var factory OperatorFactoryFace
	factory = PlusOperatorFactory{}
	t.Log(compute(factory, 1, 2))

	factory = MinusOperatorFactory{}
	t.Log(compute(factory, 4, 2))
}