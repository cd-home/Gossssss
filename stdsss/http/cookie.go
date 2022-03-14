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

	ck := &http.Cookie{
		Name:     "GodYao",
		Value:    "27",
		HttpOnly: true,
		MaxAge:   1 * 60 * 60 * 2,
	}
	http.SetCookie(rw, ck)
	rw.Header().Add("Content-Type", "application/json")
	data, _ := json.Marshal(rsp)
	fmt.Fprintln(rw, string(data))
}


// SetCookies Handler
func GetCookies(rw http.ResponseWriter, r *http.Request) {
	ck, _ := r.Cookie("GodYao")
	fmt.Fprintln(rw, string(ck.Name + ck.Value))
}

func CookieServer() {
	http.HandleFunc("/setcookie", SetCookies)
	http.HandleFunc("/getcookie", GetCookies)
	http.ListenAndServe(":8080", nil)
}
