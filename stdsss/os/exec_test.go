package main

import (
	"os/exec"
	"testing"
)

func TestExecShell(t *testing.T) {
	exec.Command("sh", "./test.sh")
}
