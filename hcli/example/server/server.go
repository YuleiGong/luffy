package main

import (
	"mime/multipart"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer() *http.Server {
	router := gin.Default()
	router.GET("/hello", handler)
	router.GET("/timeout", timeoutHandler)
	router.GET("/upload", uploadHandler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return srv
}

type Resp struct {
	Result string `json:"result"`
}

func handler(c *gin.Context) {
	c.JSON(http.StatusOK, Resp{
		Result: "hello world",
	})
}

func timeoutHandler(c *gin.Context) {
	time.Sleep(3 * time.Second)
	c.JSON(http.StatusOK, Resp{
		Result: "hello world",
	})
}

func uploadHandler(c *gin.Context) {
	var err error
	var header *multipart.FileHeader
	if header, err = c.FormFile("upload"); err != nil {
		c.Abort()
		return
	}

	var src multipart.File
	if src, err = header.Open(); err != nil {
		c.Abort()
		return
	}
	defer src.Close()

	data := make([]byte, 100)
	if _, err = src.Read(data); err != nil {
		c.Abort()
		return
	}
	c.String(http.StatusOK, string(data))
}

func main() {
	srv := InitServer()
	srv.ListenAndServe()
}
