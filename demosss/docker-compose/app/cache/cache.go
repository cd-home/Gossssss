package cache

import "github.com/go-redis/redis/v8"

var Rdb = redis.NewClient(&redis.Options{
	// 容器中就可以使用镜像的名字访问
	Addr:     "redis:6379",
	Password: "123456x",
	DB:       0,
})
