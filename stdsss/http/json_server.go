package http

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// EncodeJSON Handler
func EncodeJSON(rw http.ResponseWriter, r *http.Request) {
	type JSONResponse struct {
		Code int         `json:"code"`
		Msg  string      `json:"msg"`
		Data interface{} `json:"data"`
	}
	rw.Header().Set("Content-Type", "application/json")
	rsp := &JSONResponse{
		Code: 1,
		Msg:  "Hello World JSON",
		Data: "data",
	}
	data, _ := json.Marshal(rsp)
	fmt.Fprintln(rw, string(data))
}

type DecodeInputJSON struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

// EncodeJSON Handler
func DecodeJSON(rw http.ResponseWriter, r *http.Request) {
	var decodeInputJson DecodeInputJSON
	json.NewDecoder(r.Body).Decode(&decodeInputJson)
	fmt.Fprintf(rw, "%s %s", decodeInputJson.Title, decodeInputJson.Author)
}

func JSONServer() {
	// http.HandleFunc("/", HelloWorld)
	http.HandleFunc("/encode", EncodeJSON)
	http.HandleFunc("/decode", DecodeJSON)
	http.ListenAndServe(":8080", nil)
}
