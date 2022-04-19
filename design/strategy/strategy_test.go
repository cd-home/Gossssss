package strategy

import (
	"Gossssss/design/strategy/algs"
	"testing"
)

func TestStrategy(t *testing.T) {
	// Client 客户端调用
	lfu := &algs.Lfu{}
	cache := algs.InitCache(lfu)

	cache.Add("a", "1")
	cache.Add("b", "2")
	cache.Add("c", "3")

	lru := &algs.Lru{}
	// 替换算法
	cache.SetEvictionAlgo(lru)

	cache.Add("d", "4")

	fifo := &algs.Fifo{}
	cache.SetEvictionAlgo(fifo)
	cache.Add("e", "5")
}
