package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
func main() {
	router.GET("/ping", Ping)
	router.Run(":8081")
}
