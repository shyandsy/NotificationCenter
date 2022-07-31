package sdk

import (
	pb "github.com/shyandsy/notification-center/internal/pb"
)

type NotificationClientSdk interface {
	Start(addr string) error
	Subscribe(topic pb.Topic, callback func(string) error) error
}

type Config struct{}
type ServerEventHandler interface {
	OnConnected()
	OnDisconnected()
	OnReceive(payload []byte)
}

type ClientEventHandler interface {
	OnConnected()
	OnDisconnected()
	OnReceive(payload []byte)
}

type TopicWrapper struct {
	pb.Topic
}
