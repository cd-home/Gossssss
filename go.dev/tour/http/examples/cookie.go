package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// SetCookies Handler
func SetCookies(rw http.ResponseWriter, r *http.Request) {
	type JSONResponse struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	rsp := &JSONResponse{
		Code: 1,
		Msg:  "Hello World JSON",
		Data: "data",
	}
	// Name must not include spaces
	ck := &http.Cookie{
		Name:     "GodYao",
		Value:    "27",
		HttpOnly: true,
		MaxAge:   1 * 60 * 60 * 2,
	}
	// Set Cookie
	http.SetCookie(rw, ck)

	// Set Header
	rw.Header().Add("Content-Type", "application/json")

	// Set HTTP Code
	rw.WriteHeader(http.StatusOK)
	data, _ := json.Marshal(rsp)
	fmt.Fprintln(rw, string(data))
}

// SetCookies Handler
func GetCookies(rw http.ResponseWriter, r *http.Request) {
	ck, _ := r.Cookie("GodYao")

	rw.Header().Set("Location", "https://www.baidu.com")
	// rw.WriteHeader(http.StatusMovedPermanently) // 永久重定向
	rw.WriteHeader(http.StatusFound) // 临时重定向
	fmt.Fprintln(rw, string(ck.Name+ck.Value))
}

func CookieServer() {
	http.HandleFunc("/setcookie", SetCookies)
	http.HandleFunc("/getcookie", GetCookies)
	http.ListenAndServe(":8080", nil)
}
