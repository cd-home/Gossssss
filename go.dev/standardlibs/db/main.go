package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// open 不是真正的去链接数据库， db本身就是自带连接池, 并且默认使用
	// 所以不要去多次Open和Close
	db, err := sql.Open("mysql", "root:root@123456@tcpdemo(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	// 连接数据库
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	// 打开最大链接  (连接数 = ((核心数 * 2) + 有效磁盘数))， == Open + Idle
	// 避免 "too many connections"
	// show variables like 'max_connections' 16 + 64
	db.SetMaxOpenConns(32)

	// 默认2个，设置最大空闲链接数（直接初始化）
	db.SetMaxIdleConns(18)

	// 链接不活动最大生存时间60秒, 链接达到最大60s就会关闭； 连接被使用后，空闲时长会被重置
	db.SetConnMaxIdleTime(60)

	// 链接最长存活时间，创建开始计算（MySQL是认为链接的idle时间超过8小时就会关闭）
	// 通常情况下，该值必须比 MySQL的wait_timeout 小
	// db.SetConnMaxLifetime()
}
