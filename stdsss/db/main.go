package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func main() {
	// open 不是真正的去链接数据库， db本身就是自带连接池
	// 所以不要去多次Open和Close
	db, err := sql.Open("mysql", "root:root@123456@tcp(127.0.0.1:3306)/hello")
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

	//var name sql.NullString

	// 执行
	//result, err := DB.Exec()
	//result.LastInsertId()
	//result.RowsAffected()

	// 预处理sql， 将sql预处理为stmt，后面可以多次使用，也可以防止sql注入
	//stmt, err := DB.Prepare()
	//stmt.Exec()
	//stmt.Close()
	// 查询一条
	//row := DB.QueryRow()
	//row.Scan()

	// 查询多条
	//data := make([]map[string]xxstruct{}, 0)
	//data := make([]map[string]interface{}, 0)

	/*rows, err := DB.Query()
	rows.Scan()
	defer rows.Close()    // 每次 db.Query 操作后，都建议调⽤用 rows.Close()
	for rows.Next() {
		rows.Scan()
	 //.....  扫描出来的可以以结构体追加到切片， 或者扫描为map追加到切片
	 //data = append(data, temp)
	}*/


	// 事务
	/*tx, err := DB.Begin()
	result1, err := tx.Exec()
	result2, err := tx.Exec()
	rowsAffected1, err := result1.RowsAffected()
	rowsAffected2, err := result2.RowsAffected()
	// 链接在commit或者rollback释放
	if rowsAffected1 == 1 && rowsAffected2 == 1 {
		_ = tx.Commit()
	} else {
		_ = tx.Rollback()
	}*/

	// 推荐使用 sqlx 对数据的获取相对简单些
}
