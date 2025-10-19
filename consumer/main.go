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
	groupID    = "demo-group"
	maxRetries = 5
	retryDelay = 2 * time.Second
)

func main() {
	// 建立 Kafka reader
	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{kafkaURL},
		Topic:    topicName,
		GroupID:  groupID,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	defer reader.Close()

	fmt.Println("Consumer 已啟動，等待訊息...")

	// 持續監聽訊息
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

		message, err := reader.ReadMessage(ctx)
		cancel()

		if err != nil {
			log.Printf("讀取失敗: %v", err)
			time.Sleep(retryDelay)
			continue
		}

		// 成功讀取訊息
		fmt.Printf("[Consumer] Received message: %s\n", string(message.Value))
		fmt.Printf("[Consumer] Processed successfully\n")
	}
}
