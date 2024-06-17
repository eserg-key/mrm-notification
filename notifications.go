package mrm_notification

import (
	"github.com/eserg-key/mrm-notification/model"
	"github.com/nats-io/nats.go"
	"log"
)

const (
	typeClientDev  = "dev"
	typeClientProd = "prod"
)

type ClientInterface interface {
	Close()
	ProjectPublish(project model.Project) error
	NotificationPublish(notifications model.NotificationsInput) error
}

type Client struct {
	conn              *nats.Conn        // Соединение
	encodedConn       *nats.EncodedConn // Формат общения с nats
	topicProject      string            // Топик для регистрации проекта
	topicNotification string            // Топик для уведомления
}

func NewNotificationClient(typeClient, url string) *Client {
	var config ClientConfig
	switch typeClient {
	case typeClientDev:
		config = &DevClientConfig{}
	case typeClientProd:
		config = &ProdClientConfig{}
	default:
		log.Fatalf("notification type selection error")
	}

	c, err := nats.Connect(url)
	if err != nil {
		log.Fatalf("nats connection fail. err= %v", err)
	}

	ec, err := nats.NewEncodedConn(c, nats.JSON_ENCODER)
	if err != nil {
		log.Fatalf("nats encoded connection fail. err= %v", err)
	}

	return &Client{
		conn:              c,
		encodedConn:       ec,
		topicProject:      config.getProjectTopic(),
		topicNotification: config.getNotificationTopic(),
	}
}

func (n *Client) Close() {
	n.conn.Close()
	n.encodedConn.Close()
}

func (n *Client) ProjectPublish(project model.Project) error {
	return n.encodedConn.Publish(n.topicProject, project)
}

func (n *Client) NotificationPublish(notifications model.NotificationsInput) error {
	return n.encodedConn.Publish(n.topicNotification, notifications)
}
