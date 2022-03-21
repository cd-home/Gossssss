package main

import (
	"fmt"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("./configs/casbin_rabc_rule.conf", "./configs/casbin_policy_rule.csv")
	if err != nil {
		panic(err)
	}
	e.LoadPolicy()
	ok, _ := e.Enforce("alice", "data/1", "read")
	fmt.Println(ok)
}
