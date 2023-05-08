package main

import (
	"bytes"
	"go-web-server/wandoapp"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"
)

// file upload test
func TestUploadTest(t *testing.T) {
	path := "/Users/wando/Downloads/petbuddy.png"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	mul, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	if err != nil {
		t.Fatal("err 발생:", err)
	}

	io.Copy(mul, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	mux := wandoapp.NewHttpHandler()
	mux.ServeHTTP(res, req)
	if res.Code != http.StatusOK {
		t.Errorf("http status가 다름, 기대값: %v, 출력값: %v", http.StatusOK, res.Code)
	}

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err = os.Stat(uploadFilePath)
	if err != nil {
		t.Fatal("파일을 읽는데 에러:", err)
	}

	uploadFile, _ := os.Open(uploadFilePath)
	originFile, _ := os.Open(path)
	defer uploadFile.Close()
	defer originFile.Close()

	uploadData := []byte{}
	originData := []byte{}
	uploadFile.Read(uploadData)
	originFile.Read(originData)

	if !bytes.Equal(originData, uploadData) {
		t.Fatal("업로드된 파일과 기존 파일이 같지가 않아!")
	}
}
