package build

func Escape(num int) *int {
	 var v1, v2, v3, v4 = 1, 2, 3, 4

	 println(&num, &v1, &v2, &v3, &v4)
	// 防止内联优化
	for i := 0; i < 5; i++ {
		println(&num, &v1, &v2, &v3, &v4)
	}
	// 返回局部变量地址，runtime.newObject() 逃逸
	return &v3
}
