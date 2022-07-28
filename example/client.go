package main

import (
	"context"
	"io"
	"log"
	"sync"
	"time"

	pb "github.com/shyandsy/notification-center/pb"

	"google.golang.org/grpc/credentials/insecure"

	"google.golang.org/grpc"
)

func Start(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
		return nil, err
	}

	return conn, nil
}

func Subscribe(conn *grpc.ClientConn, topic *pb.Topic, callback func(string) error) error {
	// 初始化Greeter服务客户端
	c := pb.NewNotificationClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), 10000*time.Second)
	// 延迟关闭请求会话
	defer cancel()

	stream, err := c.Subscribe(ctx, topic)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	for {
		// 通过 Recv() 不断获取服务端send()推送的消息
		resp, err := stream.Recv()
		// 4. err==io.EOF则表示服务端关闭stream了 退出
		if err == io.EOF {
			log.Println("server closed")
			break
		}
		if err != nil {
			log.Printf("Recv error:%v", err)
			continue
		}
		if err = callback(resp.Payload); err != nil {
			log.Printf("Recv error:%v", err)
			continue
		}
	}

	return nil
}

func Callback(payload string) error {
	log.Println(payload)
	return nil
}

func main() {
	var wg sync.WaitGroup

	topic := &pb.Topic{Name: "hello_world", Actions: []string{"GET", "CREATE", "UPDATE", "DELETE"}}
	conn, err := Start("localhost:9001")
	if err != nil {
		log.Fatal(err)
	}
	// 延迟关闭连接
	defer conn.Close()

	wg.Add(1)
	go func() {
		err := Subscribe(conn, topic, Callback)
		if err != nil {
			wg.Done()
			log.Println("subscribe exit")
		}
	}()

	wg.Wait()
	log.Println("client exit")
}
