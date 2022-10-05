package rabbitmq

import (
	"context"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"go.mongodb.org/mongo-driver/bson"
)

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
}

func constructURL(config *Config) string {
	return "amqp://" + config.User + ":" + config.Password + "@" + config.Host + ":" + config.Port + "/"
}

func NewChannel(config *Config) (*amqp.Channel, error) {
	url := constructURL(config)
	conn, err := amqp.Dial(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return conn.Channel()
}

type NotificationBody struct {
	Title   string `bson:"title"`
	Message string `bson:"message"`
	Link    string `bson:"link"`
	UserID  string `bson:"userId"`
}

const NOTIFICATION_QUEUE_NAME = "notification"

func PublishNotification(ctx context.Context, channel *amqp.Channel, body *NotificationBody) error {
	bodyBytes, err := bson.Marshal(body)
	if err != nil {
		return err
	}
	return channel.PublishWithContext(ctx, "", NOTIFICATION_QUEUE_NAME, false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        bodyBytes,
	})
}
