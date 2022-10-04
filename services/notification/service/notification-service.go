package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type notificationServiceServer struct {
	db *mongo.Database
}

func NewNotificationServiceServer(db *mongo.Database) *notificationServiceServer {
	return &notificationServiceServer{db}
}

func (s *notificationServiceServer) ReadNotifications(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := s.db.Collection("notification")
	curr, err := collection.Find(ctx, bson.D{})
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	defer curr.Close(ctx)
	var results []bson.M
	if err = curr.All(ctx, &results); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, results)
}
func (s *notificationServiceServer) MarkNotificationAsRead(c *gin.Context) {
	c.String(200, "MarkNotificationAsRead")
}
