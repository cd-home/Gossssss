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
	// Form-Data Body Fields
	//request.ParseForm()
	//request.PostForm.Get()

	// Not include URL
	//request.PostFormValue()

	// URL AND Form-Data Body
	//request.ParseForm()
	//request.Form.Get()

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
	file, _ := os.Open("/Users/admin/Documents/Command_Line_Tools_for_Xcode_14.dmg")
	defer file.Close()
	w.Header().Set("Content-Disposition", `attachment; filename=`+file.Name())
	io.Copy(w, file)
}

func FileUpAndDownServer() {
	http.HandleFunc("/upload", UploadFile)
	http.HandleFunc("/dw", Download)
	http.ListenAndServe(":8080", nil)
}
