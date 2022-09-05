package fan

import "testing"

func TestFinInOut(t *testing.T) {
	ins := Producer(10)
	// 每个Square 启动一个 g 去读取 outs
	res1 := Square(ins)
	res2 := Square(ins)
	res3 := Square(ins)

	// 合并结果
	for n := range Merge(res1, res2, res3) {
		t.Logf("%3d", n)
	}
}
