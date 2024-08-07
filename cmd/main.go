package main

import (
	"context"
	"github.com/ehsaniara/delay-box/config"
	"github.com/ehsaniara/delay-box/core"
	"github.com/ehsaniara/delay-box/httpserver"
	"github.com/ehsaniara/delay-box/kafka"
	"github.com/ehsaniara/delay-box/storage"
	"github.com/ehsaniara/delay-box/worker"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	log.Printf("🛰 Starting scheduler app \n")
	ctx, cancel := context.WithCancel(context.Background())

	c := config.GetConfig()
	config.Print(c)

	// Redis localhost:6379, secret
	s, redisClientClose := storage.NewRedisClient(ctx, c)

	// RR Partition Producer
	producer, producerClose := kafka.NewProducer(c)

	// starting the scheduler app
	scheduler := core.NewScheduler(ctx, s, producer, c)
	scheduler.SetUpSubscriber(ctx)

	taskExecutor := worker.NewTaskExecutor(c, s)
	if c.WorkerEnable {
		taskExecutor.SetUpSubscriber(ctx)
	}

	//start server
	stopServer := httpserver.NewServer(ctx, nil, scheduler, c)
	log.Println("🚀 scheduler is ready!")

	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	//wait until os stop signal arrives
	keepRunning := true
	for keepRunning {
		select {
		case <-ctx.Done():
			log.Println("🖐️ terminating: context cancelled")
			keepRunning = false
		case <-sigterm:
			log.Println("🖐️ terminating: via signal")
			keepRunning = false
		}
	}
	log.Println("⏳  Stopping all services...")
	cancel()
	scheduler.Stop()
	if c.WorkerEnable {
		taskExecutor.Stop()
	}
	redisClientClose()
	producerClose()
	stopServer()

	log.Println("💔 Bye!")
}
