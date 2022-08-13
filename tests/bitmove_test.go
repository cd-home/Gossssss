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
