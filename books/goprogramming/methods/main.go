package main

import (
	"math"
)

func main() {

}

type Point struct {
	X, Y float64
}

// Distance
// Point Receiver 接收器 [简短、一致] 通常采用类型的某几个字母
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}
