package testdata

import (
	"testing"

	"github.com/casbin/casbin/v2"
)
// 用户角色 与 资源角色
func TestMultiRabcCasbin(t *testing.T) {
	e, err := casbin.NewEnforcer("../configs/rabc3/casbin_rabc_rule.conf", "../configs/rabc3/casbin_policy_rule.csv")
	if err != nil {
		panic(err)
	}

	e.LoadPolicy()

	// admin
	sub, obj, act := "admin", "data/1", "read"
	ok, _ := e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "admin", "data/1", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "admin", "data/2", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "admin", "data/2", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	// not in policy data 
	sub, obj, act = "admin", "data/3", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	// bob prod can not
	sub, obj, act = "bob", "data/1", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)
	// dev can 
	sub, obj, act = "bob", "data/2", "read"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)

	sub, obj, act = "bob", "data/2", "write"
	ok, _ = e.Enforce(sub, obj, act)
	t.Logf("sub = %s, obj = %s, act = %s, ok=%v", sub, obj, act, ok)
}
