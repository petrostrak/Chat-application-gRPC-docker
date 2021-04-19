package main

import (
	"context"
	"fmt"
	"sync"

	"github.com/petrostrak/Chat-application-gRPC-docker/proto"
)

var (
	client proto.BroadcastClient
	wait   *sync.WaitGroup
)

func init() {
	wait = &sync.WaitGroup{}
}

func connect(user *proto.User) error {
	var streamerror error
	stream, err := client.CreateStream(context.Background(), &proto.Connect{
		User:   user,
		Active: true,
	})

	if err != nil {
		return fmt.Errorf("connection failed: %v", err)
	}

	wait.Add(1)
	go func(str proto.Broadcast_CreateStreamClient) {
		defer wait.Done()
		for {
			msg, err := str.Recv()
			if err != nil {
				streamerror = fmt.Errorf("error reading message: %v", err)
				break
			}
			fmt.Printf("%v : %s\n", msg.Id, msg.Content)
		}
	}(stream)

	return streamerror
}

func main() {

}
