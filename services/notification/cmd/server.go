package cmd

import (
	"fmt"
	"log"

	"github.com/dorm-parcel-manager/dpm/common/mongo"
	"github.com/dorm-parcel-manager/dpm/common/rabbitmq"
	"github.com/dorm-parcel-manager/dpm/services/notification/config"
	"github.com/dorm-parcel-manager/dpm/services/notification/service"
	"github.com/gin-gonic/gin"
)

func RunServer() {
	configConfig := config.ProvideConfig()
	mongoDb, err := mongo.NewDb(configConfig.DB)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}
	channel, err := rabbitmq.NewChannel(configConfig.Rabbitmq)
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ")
	}
	notificationService := service.NewNotificationService(mongoDb, channel)
	go notificationService.ListenForRabbitmq()
	port := configConfig.Server.Port
	r := gin.Default()
	r.GET("/notification", notificationService.ReadNotifications)
	r.PUT("/notification", notificationService.MarkNotificationAsRead)
	err = r.Run(fmt.Sprintf(":%d", port))
	if err != nil {
		panic(err)
	}
}
