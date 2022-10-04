package cmd

import (
	"fmt"
	"github.com/dorm-parcel-manager/dpm/services/notification/config"
	"github.com/dorm-parcel-manager/dpm/services/notification/service"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	configConfig := config.ProvideConfig()
	port := configConfig.Server.Port
	r := gin.Default()
	r.GET("notification", service.ReadNotifications)
	r.PUT("notification", service.MarkNotificationAsRead)
	err := r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
