package service

import (
	"github.com/gin-gonic/gin"
)

func ReadNotifications(c *gin.Context) {
	c.String(200, "ReadNotifications")
}
func MarkNotificationAsRead(c *gin.Context) {
	c.String(200, "MarkNotificationAsRead")
}
