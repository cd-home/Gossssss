package main

import "time"

// options 配置项
type options struct {
	timeout time.Duration
	caching bool
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithTimeOut(t time.Duration) Option {
	return optionFunc(func(o *options) {
		o.timeout = t
	})
}

func WithCaching(cache bool) Option {
	return optionFunc(func(o *options) {
		o.caching = cache
	})
}

func Connect(addr string, opts ...Option) {
	// 默认值
	options := options{
		timeout: 10 * time.Second,
		caching: false,
	}
	// 额外配置可选项
	for _, o := range opts {
		o.apply(&options)
	}
}

func main() {
	Connect("addr", WithCaching(true), WithTimeOut(time.Second*6))
}
