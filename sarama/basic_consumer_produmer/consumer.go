package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	consumer, err := sarama.NewConsumer([]string{"localhost:9092"}, nil)
	//consumer, err := sarama.NewConsumer([]string{"localhost:2181"}, nil)
	if err != nil {
		log.Fatal("Consumer creation failed:", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition("test_topic", 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatal("Partition consumer creation failed:", err)
	}
	defer partitionConsumer.Close()

	for message := range partitionConsumer.Messages() {
		log.Printf("Message received: %s\n", string(message.Value))
	}
}
