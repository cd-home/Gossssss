package algs

import "fmt"

// Fifo 具体策略
type Fifo struct {
}

// Evict 策略实现
func (l *Fifo) Evict(c *Cache) {
	fmt.Println("Evicting by fifo strtegy")
}

type Lru struct {
}

func (l *Lru) Evict(c *Cache) {
	fmt.Println("Evicting by lru strtegy")
}

type Lfu struct {
}

func (l *Lfu) Evict(c *Cache) {
	fmt.Println("Evicting by lfu strtegy")
}
