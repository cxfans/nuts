package main

import (
	"encoding/json"
	"errors"
	"github.com/cxfans/nuts"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"strings"
)

const html = `<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <title>图片文字识别</title>
</head>
<body>
<form method="POST" enctype="multipart/form-data">
    请上传图片文件: <input name="image" type="file"/>
    <input type="submit" />
</form>
</body>
</html>`

const (
	apiKey    = "z9ILc5DopWA5rm4NuAou64GY"
	secretKey = "fAHDaKibDPPN8G80qTTZXxjcBA6yHYUs"
	apiUrl    = "https://aip.baidubce.com/rest/2.0/ocr/v1/accurate_basic"
)

func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		_, _ = w.Write([]byte(err.Error()))
	}
}

func uploadHandle(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()
	if r.Method != "POST" {
		_, _ = w.Write([]byte(html))
	} else {
		// 接收图片
		uploadFile, fileHeader, err := r.FormFile("image")
		errorHandle(err, w)

		// 检查图片后缀
		ext := strings.ToLower(path.Ext(fileHeader.Filename))
		if ext != ".jpg" && ext != ".png" {
			errorHandle(errors.New("only supports jpg/png"), w)
		} else {
			saveFile, err := os.Create("./uploads/" + fileHeader.Filename)
			errorHandle(err, w)
			_, _ = io.Copy(saveFile, uploadFile)

			_ = uploadFile.Close()
			_ = saveFile.Close()

			type format struct {
				Content []string `json:"content"`
			}
			var c format
			client := nuts.NewClient(apiKey, secretKey, apiUrl)
			if words, err := client.GetWordsFromImage("./uploads/" + fileHeader.Filename); err == nil {
				c.Content = words
				r, _ := json.Marshal(c)
				_, _ = w.Write(r)
			} else {
				errorHandle(err, w)
			}
		}
	}
}

func main() {
	http.HandleFunc("/", uploadHandle)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
