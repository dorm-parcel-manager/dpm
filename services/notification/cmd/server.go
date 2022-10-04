package cmd

import (
	"fmt"
	"github.com/dorm-parcel-manager/dpm/common/mongo"
	"github.com/dorm-parcel-manager/dpm/services/notification/config"
	"github.com/dorm-parcel-manager/dpm/services/notification/service"
	"github.com/gin-gonic/gin"
	"log"
)

func RunServer() {
	configConfig := config.ProvideConfig()
	mongoDb, err := mongo.NewDb(configConfig.DB)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}
	notificationServiceServer := service.NewNotificationServiceServer(mongoDb)
	port := configConfig.Server.Port
	r := gin.Default()
	r.GET("notification", notificationServiceServer.ReadNotifications)
	r.PUT("notification", notificationServiceServer.MarkNotificationAsRead)
	err = r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
