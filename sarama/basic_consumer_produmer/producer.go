package main

import (
	"log"

	"github.com/Shopify/sarama"
)

func main() {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	//producer, err := sarama.NewSyncProducer([]string{"localhost:2181"}, nil)
	if err != nil {
		log.Fatal("Producer creation failed:", err)
	}
	defer producer.Close()

	msg := &sarama.ProducerMessage{
		Topic: "test_topic",
		Value: sarama.StringEncoder("Hello Kafka"),
	}

	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatal("Message send failed:", err)
	}

	log.Printf("Message sent to partition %d with offset %d\n", partition, offset)
}
