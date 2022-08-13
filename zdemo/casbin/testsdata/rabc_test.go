package testdata

import (
	"testing"

	"github.com/casbin/casbin/v2"
)

// 用户角色
func TestRabcCasbin(t *testing.T) {
	e, err := casbin.NewEnforcer("../configs/rabc/casbin_rabc_rule.conf", "../configs/rabc/casbin_policy_rule.csv")
	if err != nil {
		panic(err)
	}

	e.LoadPolicy()

	// root
	sub, obj, act := "root", "data/1", "read"
	ok, _ := e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	// admin
	sub, obj, act = "admin", "data/1", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "admin", "data/2", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	// users
	sub, obj, act = "alice", "data/1", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "alice", "data/1", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "alice", "data/2", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "alice", "data/2", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	// role bob ==> developer == > data/2 read
	sub, obj, act = "bob", "data/1", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "bob", "data/2", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)
}
