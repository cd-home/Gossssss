package testdata

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestQueryStruct(t *testing.T) {
	t.Helper()
	db, err := sql.Open("mysql", "root:root@123456@tcpdemo(127.0.0.1:3306)/hello")
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		t.Fatal(err)
	}
	type user struct {
		id   int
		name string
	}
	// 查询所有满足条件的列，rows 应当看作是游标，而不是数据集
	rows, err := db.Query("SELECT id, name FROM user")
	if err != nil {
		t.Log(err)
	}
	// 注意必须关闭
	defer rows.Close()
	var users []user
	// 迭代获取
	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.id, &u.name)
		if err != nil {
			t.Fatal(err)
		}
		users = append(users, u)
	}
	t.Log(users)
}
