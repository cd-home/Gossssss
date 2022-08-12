package testsss

import (
	"testing"
)

/*

1 操作数在左边，  右移就是除法  m >> n   m / 2**n  2的多少次方
2.              左移就是乘法  m << n   m * 2**n

*/
func TestBitMoves(t *testing.T) {
	t.Log(1 << 10)

	t.Log(8 >> 1)
	t.Log(8 >> 2)

	t.Log(2 << 10) // 2 ** 2^10
	t.Log(67 << 1) // 67 * 2^1
}

// 编译器确定了类型
const m = uint(9)

var n = uint(9)

var a uint8 = (1 << m) / 128
var b uint8 = (1 << n) / 128

func TestBitMoves2(t *testing.T) {
	t.Log(m)
	t.Log(n)
	t.Log(a)
	t.Log(b)
}
