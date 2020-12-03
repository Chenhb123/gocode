package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/upload", upload)

	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20 // 8 MiB
	router.POST("/multi/upload", multiUpload)
	// 此处使用相对路径，实际路径根据模板所在路径作相应调整
	router.LoadHTMLGlob("../templates/*")
	router.GET("/upload", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	router.Run(":8000")
}

func upload(c *gin.Context) {
	// single file
	name := c.DefaultPostForm("name", "template")
	fmt.Println("name:", name)
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "bad request: %s", err.Error())
		return
	}
	filename := header.Filename
	out, err := os.Create(filename)
	if err != nil {
		c.String(http.StatusNotFound, "file create err: %s", err.Error())
		return
	}
	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		c.String(http.StatusNotFound, "file copy err: %s", err.Error())
		return
	}
	// c.String(http.StatusCreated, "upload successfully")
	c.Request.Header.Add("Authorization", "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1ODk4NTUyMzUsImlhdCI6MTU4OTg1NDYzNSwiaXNzIjoiZGFuYXN0dWRpb19hdXRoX3NlcnZlciIsIm1vZHVsZXMiOlsiZGV2Y2VudGVyIiwib3BlcmF0aW5nIl0sInJvbGVfaWQiOiJkZXZlbG9wZXIiLCJyb2xlX25hbWUiOiLlvIDlj5HogIUiLCJ0a3ZlcnNpb24iOjAsInVzZXJfaWQiOiJBWEV1dFZPVGlFUTdHYlMtY2FkcCIsInVzZXJfbmFtZSI6ImRldiJ9.DcJckCeNBRncQ4X1rXwxuS0lOTf1uE_4XTVcix5SGR1YhptIDO_bMDupBfwd0EOwBAYRzQxxbZCyZMuMRVguhis-8WmpXuOC17SHrqH8Vi7O6MRwwCuvOnx656jwldoOJdLw1OBpoCC5YNrTY8hYZW0fXdsqYNmV9Y6BQ7ncsx8Gzuy5G1JbnSHha7dqE6GjceHjIGBQr8JpfXFwklZV1PmBfwr0aozRIzcMGUcAxVhBItkw7oaqbWxZljd7H7N-ndohHKPfKTZiJg23Fh1UD8qW0GYkyWR5-UKkY_gZXUf3WEULAYyrUiWD7svHk8pJlK2FJTMlko8t4oUymUn2OLEBh96DlbdrnoYZdZkuPrNZDB53M-0gpMXNnj_JaM062HbIXmBNjrHmB8DG06lXSdhqyEaBzMrE-yqg4wz5VmTyY7d_60ojIqtDRfk0RhpBmiTN8PLDq6imWOWBu0QhwBj3HMz3EvPJB-wxMKeza7ZNt4KI-sy_uM_2fraiyHnHC6Zjzcs9J7VEJ32ExhxBz3VIwe8CvmoFRIcUmME0Fr1NsyKy93GPC0IAdXsQ4D6_Yv36uimurq-jAIJCsZG42MCho5dpuT4kHF90c2yywvEC_G3ZXjUR0IqlEQvaOE62r5dgdIEtci_p5F2KPalpuH1NfglYRAw6QPRHF4ZgaEc")
	c.Redirect(http.StatusPermanentRedirect, "http://192.168.80.140/danastudio/#/user/login")
}

func multiUpload(c *gin.Context) {
	// multipart form
	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		c.String(http.StatusBadRequest, "request body out of memeory: %s", err.Error())
		return
	}
	form := c.Request.MultipartForm
	files := form.File["file"]
	for i := range files {
		file, err := files[i].Open()
		if err != nil {
			c.String(http.StatusBadRequest, "file open err: %s", err.Error())
			return
		}
		defer file.Close()
		out, err := os.Create(files[i].Filename)
		if err != nil {
			c.String(http.StatusNotFound, "file create err: %s", err.Error())
			return
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			c.String(http.StatusNotFound, "file copy err: %s", err.Error())
			return
		}
	}
	c.String(http.StatusCreated, "upload successfully")
}
