package testdata

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"testing"
)

func TestQueryTransaction(t *testing.T) {
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
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	stmt, err := tx.Prepare("INSERT INTO user(id, name) VALUE(?, 'joker')")
	if err != nil {
		t.Fatal(err)
	}
	var pkErr error
	for i := 10; i <= 20; i++ {
		t.Log(i)
		if i == 20 {
			// 模拟主键冲突
			_, pkErr = stmt.Exec(1)
			t.Log(pkErr)
		}else {
			stmt.Exec(i)
		}
	}
	// 判断执行失败，全部回滚
	if pkErr != nil {
		err := tx.Rollback()
		if err != nil {
			t.Fatal(err)
		}
	} else {
		err = tx.Commit()
		if err != nil {
			t.Fatal(err)
		}
	}
	stmt.Close()
}
