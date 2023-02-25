package algs

// evictionAlgo 策略接口，声明上下文执行策略的方法
type evictionAlgo interface {
	Evict(c *Cache)
}
