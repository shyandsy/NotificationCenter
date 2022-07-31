package main

import (
	"fmt"

	"github.com/shyandsy/notification-center/internal/server"
)

const PORT = 9001

var ch chan string
var serverStop chan int
var subscribeStop chan int

type handler struct {
}

func (h handler) OnConnected()             {}
func (h handler) OnDisconnected()          {}
func (h handler) OnReceive(payload []byte) {}

func main() {
	ch = make(chan string)
	serverStop = make(chan int)
	subscribeStop = make(chan int)
	h := handler{}

	go func() {
		err := server.StartServer(PORT, h)
		if err != nil {

		}
	}()

	line := ""
	for true {
		fmt.Print("enter payload: ")
		fmt.Scanln(&line)
		ch <- line
	}
}
