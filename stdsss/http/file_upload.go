package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// UploadBigFile TODO upload
func UploadBigFile(w http.ResponseWriter, r *http.Request) {
	mr, err := r.MultipartReader()
	if err != nil {
		fmt.Sprintln(err)
		fmt.Fprintln(w, err)
		return
	}

	values := make(map[string][]string, 0)
	maxValueBytes := int64(10 << 20)
	for {
		part, err := mr.NextPart()
		if err == io.EOF {
			break
		}
		if part == nil {
			return
		}
		name := part.FormName()
		if name == "" {
			continue
		}

		fileName := part.FileName()

		fmt.Println(fileName)

		var b bytes.Buffer

		if fileName == "" {
			n, err := io.CopyN(&b, part, maxValueBytes)
			if err != nil && err != io.EOF {
				fmt.Sprintln(err)
				fmt.Fprintln(w, err)

				return
			}

			maxValueBytes -= n
			if maxValueBytes <= 0 {
				msg := "multipart message too large"
				fmt.Fprint(w, msg)
				return
			}
			values[name] = append(values[name], b.String())
		}

		dst, err := os.Create("./static/" + fileName)

		if err != nil {
			fmt.Fprint(w, err.Error())
		}

		defer dst.Close()

		for {
			buffer := make([]byte, 100000)
			cBytes, err := part.Read(buffer)
			if err == io.EOF {
				break
			}
			dst.Write(buffer[0:cBytes])
		}

	}
}
// UploadFile small file upload
func UploadFile(writer http.ResponseWriter, request *http.Request) {
	f, header, _ := request.FormFile("upload")
	defer f.Close()
	fileName := time.Now().Format("20060102150405_") + header.Filename
	dst, _ := os.Create(fileName)
	defer dst.Close()
	_, _ = io.Copy(dst, f)
	_, _ = fmt.Fprint(writer, fileName)
}

// UploadFile small files upload
func UploadFiles(writer http.ResponseWriter, request *http.Request) {
	_ = request.ParseMultipartForm(32 << 20)
	files := request.MultipartForm.File["upload"]
	l := len(files)
	for i := 0; i < l; i++ {
		fileName := time.Now().Format("20060102150405_") + files[i].Filename
		src, _ := files[i].Open()
		dst, _ := os.Create(fileName)
		_, _ = io.Copy(dst, src)
		//src.Close()
	}
	_, _ = fmt.Fprint(writer, files)
}
