package service

import (
	"context"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

type MarkNotificationAsReadRequest struct {
	Id string `json:"_id"`
}

func (s *notificationServiceServer) MarkNotificationAsRead(c *gin.Context) {
	req := MarkNotificationAsReadRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	collection := s.db.Collection("notification")
	objId, err := primitive.ObjectIDFromHex(req.Id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	filter := bson.D{{Key: "_id", Value: objId}}
	update := bson.D{
		{Key: "$set", Value: bson.D{
			{Key: "read", Value: true},
		}},
	}
	_, err = collection.UpdateOne(ctx, filter, update)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "success"})
}
