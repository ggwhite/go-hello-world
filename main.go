package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()

	e.GET("", func(c *gin.Context) {
		name := c.Query("name")
		c.String(http.StatusOK, "Hello %s", name)
	})

	e.Run(":3000")
}
