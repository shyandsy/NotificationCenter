package sdk

import pb "github.com/shyandsy/notification-center/pb"

type NotificationClientSdk interface {
	Start(addr string) error
	Subscribe(topic pb.Topic, callback func(string) error) error
}
