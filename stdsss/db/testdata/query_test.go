package testdata

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestQuery(t *testing.T) {
	t.Helper()
	db, err := sql.Open("mysql", "root:root@123456@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
	var (
		id   int
		name string
	)
	// 查询所有满足条件的列，rows 应当看作是游标，而不是数据集
	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		t.Log(err)
	}
	// always 注意必须关闭， 关闭是有可能发生错误的，可以捕获记录
	defer rows.Close()
	for rows.Next() {
		var s sql.NullString
		// 必须是指针
		err := rows.Scan(&id, &name)
		if err != nil {
			t.Log(err)
		}
		// 判断是否是空
		// if s.Valid {

		// }
		t.Log(id, name)
	}
	// 必须在迭代获取之后，检查错误（循环可能由于某种原因退出）
	err = rows.Err()
	if err != nil {
		t.Fatal(err)
	}

	// 返回单个数据
	row := db.QueryRow("SELECT id, name FROM user WHERE id = ?", 2)
	err = row.Scan(&id, &name)
	if err != nil {
		// 无数据
		if err == sql.ErrNoRows {
			t.Log(err)
		} else {
			t.Fatal(err)
		}
	}
	t.Log(id, name)
}
