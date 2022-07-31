package server

import (
	"github.com/shyandsy/notification-center/helper"
	"github.com/shyandsy/notification-center/internal/pb"
	"github.com/shyandsy/notification-center/internal/server"
)

type Subject interface {
	Subscribe(topic __.Topic, channel chan string) error
	Unsubscribe(topic string, action string, channel chan string) error
	Notify(topic string, action string, payload string)
}

type Actions struct {
	Actions map[string]__.Notification_SubscribeServer
}

type subject struct {
	receivers map[string]*server.Receiver
}

func (s *subject) Subscribe(topic __.Topic, channel chan string) error {
	receiver, ok := s.receivers[topic.Name]
	if !ok {
		s.receivers[topic.Name] = &server.Receiver{
			Actions: topic.Actions,
			Channel: channel,
		}
	} else {
		for _, action := range topic.Actions {
			if !helper.Contains(receiver.Actions, action) {
				receiver.Actions = append(receiver.Actions, action)
			}
		}
	}
	return nil
}

func (s *subject) Unsubscribe(topic string, action string, channel chan string) error {

}

func (s *subject) Notify(topic string, action string, payload string) {

}
