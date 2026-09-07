package main

import (
	"RewriteProject/infra/mailer"
	"RewriteProject/internal/config"
	"context"
	"log"
	"time"
)

func main() {
	conf := &config.MailerConfig{
		Host:     "localhost",
		Port:     1025,
		User:     "root",
		Password: "pass",
	}

	mail := mailer.Mail{
		From:    "goapp@gmail.com",
		To:      []string{"client@gmail.com", "client2@gmail.com", "client3@gmail.com"},
		Subject: "Test Email",
		Body:    "Hello, This is a test email from Go.",
	}

	mailer := mailer.NewMailer(conf)

	ctx, cancel := context.WithCancel(context.Background())
	err := mailer.StartWorker(ctx)
	if err != nil {
		log.Fatal(err)
	}
	log.Print("Mailer Worker Start")
	log.Print("Normal Traffic")

	queue1 := mailer.Enqueue(ctx, mail)
	res1 := <-queue1
	log.Printf("Result: %+v\n", res1)

	time.Sleep(1 * time.Second)

	queue2 := mailer.Enqueue(ctx, mail)
	res2 := <-queue2
	log.Printf("Result: %+v\n", res2)

	time.Sleep(1 * time.Second) // Normal Trafic

	log.Print("High Traffic")
	start := time.Now()
	for i := 0; i < 100; i++ {
		// Simulate High Trafic
		queue3 := mailer.Enqueue(ctx, mail)
		res3 := <-queue3
		log.Printf("Result: %+v\n", res3)
	}

	end := time.Since(start)
	log.Printf("High Traffic took: %s\n", end)

	log.Print("Stopping")
	cancel()
	time.Sleep(2 * time.Second)

	for i := 0; i < 10; i++ {
		// Simulate High Trafic
		queue3 := mailer.Enqueue(ctx, mail)
		res3 := <-queue3
		log.Printf("Result: %+v\n", res3)
	}

	mailer.Greating()

}
