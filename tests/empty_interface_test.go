package testsss

import "testing"

type M struct {}

func TestEmptyInterface(t *testing.T) {
	m := M{}
	pm := &m

	// Empty interface is Empty interface

	var i interface{} = m
	var j interface{} = pm

	t.Log(i)
	t.Log(j)

	// Do not use interface pointer and not necessary
	// var k *interface{} = m
	// var h *interface{} = pm

}