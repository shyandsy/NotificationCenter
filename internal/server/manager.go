package server

import (
	"errors"
	"sync"

	pb "github.com/shyandsy/notification-center/internal/pb"
)

type ClientStreamManager interface {
	AddClientStream(topic *pb.Topic, stream pb.Notification_SubscribeServer) error
	Notify(topic string, action string, payload string) error
}

type Receiver struct {
	Actions []string
	Stream  pb.Notification_SubscribeServer
}
type clientManager struct {
	table map[string]*[]*Receiver
	list  map[*pb.Notification_SubscribeServer]interface{}
	mu    sync.Mutex
}

func NewClientManager() ClientStreamManager {
	return &clientManager{
		table: make(map[string]*[]*Receiver),
	}
}

func (m *clientManager) AddClientStream(topic *pb.Topic, stream pb.Notification_SubscribeServer) error {
	if topic == nil || stream == nil {
		return errors.New("topic and stream cannot be nil")
	}
	receiver := &Receiver{
		Actions: topic.Actions,
		Stream:  stream,
	}
	m.mu.Lock()
	if receivers, ok := m.table[topic.Name]; ok {
		*receivers = append(*receivers, receiver)
	} else {
		m.table[topic.Name] = &([]*Receiver{receiver})
	}
	m.mu.Unlock()
	return nil
}

func (m *clientManager) Notify(topic string, action string, payload string) error {
	m.mu.Lock()
	if receivers, ok := m.table[topic]; ok {
		for _, receiver := range *receivers {
			for _, item := range receiver.Actions {
				if action == item {
					err := receiver.Stream.Send(&pb.Response{Payload: payload})
					if err != nil {
						return err
					}
					break
				}
			}
		}
	}
	m.mu.Unlock()
	return nil
}
