package server

import (
	"fmt"

	"github.com/dorm-parcel-manager/dpm/services/notification/service"
	"github.com/gin-gonic/gin"
)

type Server struct {
	NotificationService *service.NotificationService
	Config              *Config
	Router              *gin.Engine
}

type Config struct {
	Port string
}

func NewServer(notificationService *service.NotificationService, config *Config) *Server {
	router := gin.Default()
	router.GET("/notification", notificationService.GetNotifications)
	router.PATCH("/notification/:id", notificationService.PatchNotificationRead)
	return &Server{
		notificationService, config, router,
	}
}

func (s *Server) Start() error {
	go s.NotificationService.ListenForRabbitmq()
	return s.Router.Run(fmt.Sprintf(":%s", s.Config.Port))
}
