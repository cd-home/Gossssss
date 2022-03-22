package testdata

import (
	"strings"
	"testing"

	"github.com/casbin/casbin/v2"
)

func KeyMatch(key1, key2 string) bool {
	if i := strings.Index(key2, "*"); i != -1 {
		return key1[:i] == key2[:i]
	}
	return key1 == key2
}
func KeyMatchFunc(args ...interface{}) (interface{}, error) {
	key1 := args[0].(string)
	key2 := args[1].(string)
	return (bool)(KeyMatch(key1, key2)), nil
}

func TestRabcCasbinDeFineKeyFunc(t *testing.T) {
	e, err := casbin.NewEnforcer("../configs/rabc2/casbin_rabc_rule.conf", "../configs/rabc2/casbin_policy_rule.csv")
	if err != nil {
		panic(err)
	}

	// define a match rule for obj
	e.AddFunction("KeyMatchFunc", KeyMatchFunc)

	e.LoadPolicy()
	// admin
	sub, obj, act := "admin", "data/*", "read"
	ok, _ := e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	// users
	sub, obj, act = "alice", "data/1", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "alice", "data/1", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	// role
	sub, obj, act = "bob", "data/1", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "bob", "data/2", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)
}
