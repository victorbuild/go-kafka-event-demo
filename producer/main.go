package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	kafkaURL   = "localhost:9092"
	topicName  = "demo-topic"
	message    = "user_registered"
	maxRetries = 5
	retryDelay = 2 * time.Second
)

func main() {
	// 建立 Kafka writer
	writer := &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topicName,
		Balancer: &kafka.LeastBytes{},
	}
	defer writer.Close()

	// 重試機制
	for i := 0; i < maxRetries; i++ {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		err := writer.WriteMessages(ctx, kafka.Message{
			Key:   []byte("demo-key"),
			Value: []byte(message),
		})

		cancel()

		if err != nil {
			if i == maxRetries-1 {
				log.Fatalf("發送失敗，已重試 %d 次: %v", maxRetries, err)
			}
			log.Printf("發送失敗，%v 後重試... (第 %d/%d 次)", retryDelay, i+1, maxRetries)
			time.Sleep(retryDelay)
			continue
		}

		// 成功發送
		fmt.Printf("[Producer] Sent message: %s\n", message)
		return
	}
}
