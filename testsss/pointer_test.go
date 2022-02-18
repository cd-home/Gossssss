package testsss

import "testing"

func TestPointer(t *testing.T) {
	var i = 0

	// pointer can modify the value
	func(j *int) {
		(*j)++
		t.Log(i)
	}(&i)

	func(j int) {
		j++
		t.Log(i)
	}(i)

	t.Log(i)
}