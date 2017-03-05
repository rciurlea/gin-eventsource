package main

import (
	"io"
	"net/http"
	"os"
	"path"
	"time"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "/static/")
	})

	cwd, _ := os.Getwd()
	r.Static("/static", path.Join(cwd, "static"))

	r.GET("/events", eventsHandler)

	r.Run("127.0.0.1:3000")
}

func eventsHandler(c *gin.Context) {
	c.Stream(func(w io.Writer) bool {
		time.Sleep(time.Second)
		c.SSEvent("ping", time.Now().String())
		return true
	})
}
