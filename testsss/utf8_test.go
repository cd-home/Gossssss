package testsss

import "testing"


func TestUTF8(t *testing.T) {
	t.Log(len("A"))
	t.Log(len("你"))
	t.Log(len("你好"))
}