package testdata

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"testing"
)

func TestInsert(t *testing.T) {
	t.Helper()
	db, err := sql.Open("mysql", "root:root@123456@tcp(127.0.0.1:3306)/hello")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	// 非返回rows ，一律使用Exec
	result, err := db.Exec(`INSERT INTO user(name) VALUES("jack"), ("mary"), ("bob")`)
	if err != nil {
		t.Fatal(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		t.Fatal(err)
	}
	t.Log(id)
}
