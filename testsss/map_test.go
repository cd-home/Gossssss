package testsss

import (
	"testing"
)

func TestNilMap(t *testing.T) {
	var testMap map[string]string
	// map cannot be nil
	// testMap["name"] = "foo"
	t.Log(testMap)

	testMap = make(map[string]string)
	testMap["name"] = "foo"
	t.Log(testMap)
}

func TestNoExistKey(t *testing.T) {
	testMap := make(map[string]string)

	// NoExist Key will Not panic (support now)
	// Print default type value
	t.Log(testMap["NoExist"])

	// ok mode
	if name, ok := testMap["NoExist"]; ok {
		t.Log(name)
	} else {
		t.Log("NoExist")
	}
}

type example struct {
	Name string
}

func TestModifyMapValue(t *testing.T) {
	var m = map[string]*example{
		"one": {Name: "GodYao"},
	}
	m["one"].Name = "Mike" // compiler error
	t.Log(m)
}

func TestModifyMapValue2(t *testing.T) {
	var m = map[string]example{
		"one": {Name: "GodYao"},
	}
	m["one"] = example{Name: "Mike"} // compiler error
	t.Log(m)
}

func TestNewFunc(t *testing.T) {
	var p *[5]int = new([5]int)
	t.Log(p)

	var pp []int = make([]int, 6)
	t.Log(pp)
}

type exampleAddr struct {
	Name string
}

func TestGetAddress(t *testing.T) {
	var a int = 10
	t.Log(&a)

	var e = &exampleAddr{Name: "Mike"}
	t.Log(&e.Name)

	var p *[5]int = new([5]int)
	t.Log(&p[0])

	var pp []int = make([]int, 6)
	t.Log(&pp[1])

	t.Log(&*p)
}

type SafeMap struct {
	// sync.Mutex
	m map[string]string
}

func (m *SafeMap) Set(key string, value string) {
	// defer m.Unlock()
	// m.Lock()
	m.m[key] = value
}

func TestSafeMap(t *testing.T) {
	mp := &SafeMap{m: make(map[string]string)}
	mp.Set("foo", "bar")
	t.Log(mp)
}
