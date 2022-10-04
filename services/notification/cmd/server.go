package cmd

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RunServer() {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Notification Service")
	})
	r.Run(":4567")
}
