package server

import (
	"fmt"
	"log"
	"net"
	"time"

	__ "github.com/shyandsy/notification-center/internal/pb"
	"github.com/shyandsy/notification-center/sdk"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

var ch chan string
var serverStop chan int
var subscribeStop chan int

var kaep = keepalive.EnforcementPolicy{
	MinTime:             5 * time.Second, // If a client pings more than once every 5 seconds, terminate the connection
	PermitWithoutStream: true,            // Allow pings even when there are no active streams
}

var kasp = keepalive.ServerParameters{
	MaxConnectionIdle:     15 * time.Second, // If a client is idle for 15 seconds, send a GOAWAY
	MaxConnectionAge:      30 * time.Second, // If any connection is alive for more than 30 seconds, send a GOAWAY
	MaxConnectionAgeGrace: 5 * time.Second,  // Allow 5 seconds for pending RPCs to complete before forcibly closing connections
	Time:                  5 * time.Second,  // Ping the client if it is idle for 5 seconds to ensure the connection is still active
	Timeout:               1 * time.Second,  // Wait 1 second for the ping ack before assuming the connection is dead
}

type notificationServer struct{}

func NewNotificationServer() {
	notificationServer{}
}

func (n *notificationServer) Subscribe(req *__.Topic, stream __.Notification_SubscribeServer) error {
	log.Printf("Recved Topic%v, %v", req.Name, req.Actions)

	// 具体返回多少个response根据业务逻辑调整
	for true {
		payload, ok := <-ch
		if !ok {
			log.Printf("ch关闭，subscribe退出")
			break
		}
		if payload == "" {
			log.Printf("输入为空，subscribe退出")
			break
		}
		// 通过 send 方法不断推送数据
		err := stream.Send(&__.Response{Payload: payload})
		if err != nil {
			log.Fatalf("Send error:%v", err)
			return err
		}
	}
	// 返回nil表示已经完成响应
	subscribeStop <- 1
	return nil
}

func StartServer(port int, handler sdk.ServerEventHandler) error {
	addr := fmt.Sprintf("127.0.0.1:%d", port)
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer(grpc.KeepaliveEnforcementPolicy(kaep), grpc.KeepaliveParams(kasp))
	srv := &notificationServer{}
	__.RegisterNotificationServer(s, srv)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	return nil
}
