package server

import (
	"context"
	"errors"
	"testing"

	pb "github.com/shyandsy/notification-center/internal/pb"
	"google.golang.org/grpc"
)

func makeStreamMock() *StreamMock {
	return &StreamMock{
		ctx:            context.Background(),
		sentFromServer: make(chan *pb.Response, 10),
	}
}

type StreamMock struct {
	grpc.ServerStream
	ctx            context.Context
	sentFromServer chan *pb.Response
}

func (m *StreamMock) Context() context.Context {
	return m.ctx
}
func (m *StreamMock) Send(resp *pb.Response) error {
	m.sentFromServer <- resp
	return nil
}
func (m *StreamMock) ReceiveToClient() (*pb.Response, error) {
	select {
	case response := <-m.sentFromServer:
		return response, nil
	default:
		return nil, errors.New("empty")
	}
}

var topics = []string{"A", "B"}
var actions = []string{"CREATE", "UPDATE", "GET", "DELETE"}

func TestManager(t *testing.T) {
	// create clients
	client1 := makeStreamMock()
	client2 := makeStreamMock()
	client3 := makeStreamMock()

	manager := NewClientManager()

	// client 1 subscribes topic A, action CREATE, UPDATE, DELETE
	if err := manager.AddClientStream(&pb.Topic{
		Name: "A", Actions: []string{"CREATE", "UPDATE", "DELETE"},
	}, client1); err != nil {
		t.Error("client stream failed")
	}

	// client 2 subscribes topic A, action CREATE
	if err := manager.AddClientStream(&pb.Topic{
		Name: "A", Actions: []string{"CREATE"},
	}, client2); err != nil {
		t.Error("client stream failed")
	}

	// client 3 subscribes topic B, action DELETE
	if err := manager.AddClientStream(&pb.Topic{
		Name: "B", Actions: []string{"DELETE"},
	}, client3); err != nil {
		t.Error("client stream failed")
	}

	// subscribe different topics
	if err := manager.Notify("A", "CREATE", "xxxxxxx"); err != nil {
		t.Error("manager notify failed")
	}
	if err := manager.Notify("A", "UPDATE", "yyyyyyy"); err != nil {
		t.Error("manager notify failed")
	}
	if err := manager.Notify("B", "DELETE", "zzzzzzz"); err != nil {
		t.Error("manager notify failed")
	}
	if err := manager.Notify("B", "CREATE", "....."); err != nil {
		t.Error("manager notify failed")
	}

	// verify
	if resp, err := client1.ReceiveToClient(); err != nil || resp.Payload != "xxxxxxx" {
		t.Error("client 1 receive nothing")
	}
	if resp, err := client2.ReceiveToClient(); err != nil || resp.Payload != "xxxxxxx" {
		t.Error("client 2 receive nothing")
	}
	if resp, err := client1.ReceiveToClient(); err != nil || resp.Payload != "yyyyyyy" {
		t.Error("client 1 receive nothing")
	}
	if _, err := client2.ReceiveToClient(); err == nil {
		t.Error("client 2 should sent nothing")
	}
	if resp, err := client3.ReceiveToClient(); err != nil || resp.Payload != "zzzzzzz" {
		t.Error("client 3 receive nothing")
	}

	if _, err := client1.ReceiveToClient(); err == nil {
		t.Error("client 1 should sent nothing")
	}
	if _, err := client2.ReceiveToClient(); err == nil {
		t.Error("client 2 should sent nothing")
	}
	if _, err := client3.ReceiveToClient(); err == nil {
		t.Error("client 3 should sent nothing")
	}
}
