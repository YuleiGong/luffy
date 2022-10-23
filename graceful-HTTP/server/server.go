package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitServer() *http.Server {
	router := gin.Default()
	router.GET("/", handler)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	return srv
}

func handler(c *gin.Context) {
	time.Sleep(5 * time.Second)
	c.String(http.StatusOK, "Welcome Gin Server")
}
