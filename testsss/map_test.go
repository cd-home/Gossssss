package testsss

import "testing"

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
