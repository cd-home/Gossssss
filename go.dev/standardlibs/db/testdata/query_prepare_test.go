package testdata

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestQueryPrepareStruct(t *testing.T) {
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

	// 预先编译sql，可以使用多次
	stmt, err := db.Prepare("SELECT id, name FROM user WHERE id > ?")
	if err != nil {
		t.Log(err)
	}
	// 注意必须关闭
	defer stmt.Close()

	// 使用多次，动态参数
	rows, err := stmt.Query(2)
	if err != nil {
		t.Fatal(err)
	}
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
