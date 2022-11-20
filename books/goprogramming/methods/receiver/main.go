package main

func main() {
	p := Point{
		X: 3,
		Y: 4,
	}
	// 调用会拷贝 receiver
	// 当 receiver 比较大时采用指针模式较为合适, 或者当需要在方法中修改其本身时
	p.ScaleBy(2)

	// 还需要说明一点, 不管是 pointer receiver 还是 value receiver
	// 调用都不必显式的做类型转换,编译器可自动转换
	p.ScaleBy(2)
	(&p).ScaleBy(2)

	// 总结
	// 不管你的method的receiver是指针类型还是非指针类型, 都是可以通过指针/非指针类型进行调用的，编译器会帮你做类型转换
	// 在声明一个method的receiver该是指针还是非指针类型时，你需要考虑两方面的因素，
	// 1. 调用会产生拷贝, 拷贝值还是指针你需要衡量
	// 2. 如果是指针接收者, 那么就可以修改原类型, 其是否符合你的需求
}

type Point struct {
	X, Y float64
}

// ScaleBy
// 一般如果有指针作为 receiver 那么其所有的方法能应该统一为 pointer receiver
func (p *Point) ScaleBy(factor float64) {
	// p 指针可直接访问
	p.X *= factor
	p.Y *= factor
}
