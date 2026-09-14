package main

import (
	sse "RewriteProject/infra/SSE"
	"context"
	"fmt"
	"log"
	"time"
)

func main() {
	manager := sse.NewUsecase()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn1 := manager.Connect(ctx, "user1")
	conn12 := manager.Connect(ctx, "user1")
	conn13 := manager.Connect(ctx, "user1")

	go func() {
		for {
			select {
			case reply := <-conn1.Reply:
				fmt.Println("conn1" + string(reply))
			case <-ctx.Done():
				manager.Disconnect(ctx, "user1", conn1)
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case reply := <-conn12.Reply:
				fmt.Println("conn2" + string(reply))
			case <-ctx.Done():
				manager.Disconnect(ctx, "user1", conn12)
				return
			}
		}
	}()
	go func() {
		for {
			select {
			case reply := <-conn13.Reply:
				fmt.Println("conn3" + string(reply))
			case <-ctx.Done():
				manager.Disconnect(ctx, "user1", conn13)
				return
			}
		}
	}()

	time.Sleep(100 * time.Millisecond)

	log.Print("Start Sending")

	manager.SendMessage(context.Background(), "user1", []byte("hello"))
	manager.SendMessage(context.Background(), "user1", []byte("am"))
	manager.SendMessage(context.Background(), "user1", []byte("From"))
	manager.Disconnect(ctx, "user1", conn1)
	manager.Disconnect(ctx, "user1", conn12)
	manager.Disconnect(ctx, "user1", conn13)
	manager.SendMessage(context.Background(), "user1", []byte("e"))
	log.Print("Done Sending")
	log.Print(manager)

	time.Sleep(10 * time.Second) // atau cara lain nunggu, bukan wg.Wait()
	cancel()

}
