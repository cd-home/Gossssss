package testdata

import (
	"testing"

	"github.com/casbin/casbin/v2"
)

func TestRabcCasbinDeFineKeyFunc(t *testing.T) {
	e, err := casbin.NewEnforcer("../configs/rabc2/casbin_rabc_rule.conf", "../configs/rabc2/casbin_policy_rule.csv")
	if err != nil {
		panic(err)
	}

	e.LoadPolicy()
	// admin
	sub, obj, act := "admin", "data/*", "read"
	ok, _ := e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "admin", "data/1", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	// mike inherit the admin
	sub, obj, act = "mike", "data/1", "read"
	ok, _ = e.Enforce(sub, obj, act)
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
