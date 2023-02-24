package main

import (
	"net/http"
	"text/template"
)

func Template(rw http.ResponseWriter, r *http.Request) {
	type Data struct {
		Name string `json:"name"`
		Age  uint8  `json:"age"`
	}
	data := Data{
		Name: "mike",
		Age:  20,
	}
	//template.HTMLEscaper()
	// 解析文件
	// tmpl := template.Must(template.ParseFiles("layout.html"))
	t, _ := template.ParseFiles("./static/index.html")
	// 填充模版
	_ = t.Execute(rw, data)
}

func main() {
	http.HandleFunc("/template", Template)
	_ = http.ListenAndServe(":8080", nil)
}
