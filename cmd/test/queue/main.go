package main

import (
	"RewriteProject/infra/queue"
	"context"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

func main() {
	rng := rand.New(rand.NewSource(1)) // seed tetap agar reproducible
	const invariantCheckProb = 0.2
	ctx, cancel := context.WithCancel(context.Background())
	Queue := queue.NewQueue[string, string, bool]("Test", 3, 1)
	Queue.Start(ctx)
	go func() {
		for process := range Queue.Job {
			process.Reply <- true
			close(process.Reply)
		}
	}()
	var wg sync.WaitGroup
	start := time.Now()
	for i := range 100 {
		key := fmt.Sprintf("job%d", i)
		job := fmt.Sprintf("job%d", i)
		if rng.Float64() < invariantCheckProb {
			wg.Add(2)
			go func(key string, job string) {
				defer wg.Done()

				res, err := Queue.Push(ctx, key, job)
				log.Print("dup: ", res, err)
			}(key, job)
			go func(key string, job string) {
				defer wg.Done()

				res, err := Queue.Push(ctx, key, job)
				log.Print("dup: ", res, err)
			}(key, job)
		}

		wg.Add(1)
		go func(key string, job string) {
			defer wg.Done()

			res, err := Queue.Push(ctx, key, job)
			log.Print(res, err)
		}(key, job)

	}

	wg.Wait()
	end := time.Since(start)
	log.Println(end)

	Queue.Close()

	cancel()

}
