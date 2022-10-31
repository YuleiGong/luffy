package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitServer() *http.Server {
	router := gin.Default()
	router.GET("/hello", handler)

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

func main() {
	srv := InitServer()
	srv.ListenAndServe()
}
