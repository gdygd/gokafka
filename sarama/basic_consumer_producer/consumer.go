package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	var consum sarama.Consumer
	var err error = nil
	//consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	consum, err = sarama.NewConsumer([]string{"localhost:9092"}, nil)
	//consumer, err := sarama.NewConsumer([]string{"localhost:2181"}, nil)
	if err != nil {
		log.Fatal("Consumer creation failed:", err)
	}
	defer consum.Close()

	//partitionConsumer, err := consumer.ConsumePartition("test_topic", 0, sarama.OffsetNewest)
	partitionConsumer, err := consum.ConsumePartition("test_topic", 0, 0)
	if err != nil {
		log.Fatal("Partition consumer creation failed:", err)
	}
	defer partitionConsumer.Close()

	go func() {
		for err := range partitionConsumer.Errors() {
			log.Printf("에러 발생: %v", err)
		}
	}()

	for message := range partitionConsumer.Messages() {
		log.Printf("Message received: %s\n", string(message.Value))
	}
}
