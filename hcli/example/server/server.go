package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer() *http.Server {
	router := gin.Default()
	router.GET("/hello", handler)
	router.GET("/timeout", timeoutHandler)

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

func main() {
	srv := InitServer()
	srv.ListenAndServe()
}
