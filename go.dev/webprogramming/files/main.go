package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// UploadFile small file upload
func UploadFile(writer http.ResponseWriter, request *http.Request) {
	// Use this OK !
	//request.FormValue()
	val := request.FormValue("name")
	fmt.Println(val)
	// One
	//request.ParseMultipartForm()
	//request.MultipartForm.File[""]
	// Two
	f, header, _ := request.FormFile("upload")
	defer f.Close()
	fileName := time.Now().Format("20060102150405_") + header.Filename
	dst, _ := os.Create(fileName)
	defer dst.Close()
	_, _ = io.Copy(dst, f)
	_, _ = fmt.Fprint(writer, fileName)
}

// UploadFiles files
func UploadFiles(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseMultipartForm(32 << 20)
	files := request.MultipartForm.File["upload"]
	l := len(files)
	for i := 0; i < l; i++ {
		fileName := time.Now().Format("20060102150405_") + files[i].Filename
		src, _ := files[i].Open()
		dst, _ := os.Create(fileName)
		_, _ = io.Copy(dst, src)
		src.Close()
	}
	_, _ = fmt.Fprint(writer, files)
}

// Download file
func Download(w http.ResponseWriter, r *http.Request) {
	path, _ := os.Getwd()
	file, _ := os.Open(path + "/go.dev/tour/webprogramming/files/main.go")
	defer file.Close()
	w.Header().Set("Content-Disposition", `attachment; filename=`+file.Name())
	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/upload", UploadFile)
	http.HandleFunc("/dw", Download)
	_ = http.ListenAndServe(":8080", nil)
}
