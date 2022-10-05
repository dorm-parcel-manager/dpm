package service

import (
	"context"
	"log"
	"time"

	"github.com/dorm-parcel-manager/dpm/common/rabbitmq"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type notificationService struct {
	db              *mongo.Database
	rabbitmqChannel *amqp.Channel
}

func NewNotificationService(db *mongo.Database, rabbitmqChannel *amqp.Channel) *notificationService {
	return &notificationService{db, rabbitmqChannel}
}

func (s *notificationService) ReadNotifications(c *gin.Context) {
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

func (s *notificationService) MarkNotificationAsRead(c *gin.Context) {
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

func (s *notificationService) ListenForRabbitmq() {
	q, err := s.rabbitmqChannel.QueueDeclare(
		rabbitmq.NOTIFICATION_QUEUE_NAME, // name
		false,                            // durable
		false,                            // delete when unused
		false,                            // exclusive
		false,                            // no-wait
		nil,                              // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %s", err)
	}
	msgs, err := s.rabbitmqChannel.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %s", err)
	}
	forever := make(chan bool)

	go func() {
		for d := range msgs {
			notificationBody := rabbitmq.NotificationBody{}
			err := bson.Unmarshal(d.Body, &notificationBody)
			if err != nil {
				log.Printf("Failed to unmarshal notification: %s", err)
				continue
			}
			s.db.Collection("notification").InsertOne(context.Background(), bson.M{
				"title":   notificationBody.Title,
				"message": notificationBody.Message,
				"link":    notificationBody.Link,
				"read":    false,
			})
		}
	}()

	<-forever
}
